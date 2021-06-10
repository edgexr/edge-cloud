package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/coreos/etcd/clientv3/concurrency"
	"github.com/mobiledgex/edge-cloud/cloud-resource-manager/platform/pc"
	"github.com/mobiledgex/edge-cloud/cloudcommon"
	"github.com/mobiledgex/edge-cloud/edgeproto"
	"github.com/mobiledgex/edge-cloud/gcs"
	"github.com/mobiledgex/edge-cloud/log"
)

type GPUDriverApi struct {
	sync  *Sync
	store edgeproto.GPUDriverStore
	cache edgeproto.GPUDriverCache
}

var gpuDriverApi = GPUDriverApi{}

const (
	GPUDriverBuildURLValidity = 20 * time.Minute
	ChangeInProgress          = "ChangeInProgress"
)

func InitGPUDriverApi(sync *Sync) {
	gpuDriverApi.sync = sync
	gpuDriverApi.store = edgeproto.NewGPUDriverStore(sync.store)
	edgeproto.InitGPUDriverCache(&gpuDriverApi.cache)
	sync.RegisterCache(&gpuDriverApi.cache)
}

// Must call GCSClient.Close() when done
func getGCSStorageClient(ctx context.Context) (*gcs.GCSClient, error) {
	bucketName := cloudcommon.GetGPUDriverBucketName(nodeMgr.DeploymentTag)
	credsObj, err := gcs.GetGCSCreds(ctx, vaultConfig)
	if err != nil {
		return nil, err
	}
	storageClient, err := gcs.NewClient(ctx, credsObj, bucketName, gcs.LongTimeout)
	if err != nil {
		return nil, fmt.Errorf("Unable to setup GCS client: %v", err)
	}
	return storageClient, nil
}

func setupGPUDriver(ctx context.Context, storageClient *gcs.GCSClient, driverKey *edgeproto.GPUDriverKey, build *edgeproto.GPUDriverBuild, cb edgeproto.GPUDriverApi_CreateGPUDriverServer) error {
	if build.DriverPath == "" {
		return fmt.Errorf("Missing driverpath: %s", build.Name)
	}
	if build.OperatingSystem == edgeproto.OSType_LINUX && build.KernelVersion == "" {
		return fmt.Errorf("Kernel version is required for Linux build %s", build.Name)
	}
	driverFileName, err := cloudcommon.GetFileNameWithExt(build.DriverPath)
	if err != nil {
		return err
	}
	ext := filepath.Ext(driverFileName)
	// Download the package
	var fileContents string
	authApi := &cloudcommon.VaultRegistryAuthApi{
		VaultConfig: vaultConfig,
	}
	cb.Send(&edgeproto.Result{Message: "Downloading GPU driver build " + build.Name})
	err = cloudcommon.DownloadFile(ctx, authApi, build.DriverPath, build.DriverPathCreds, "", &fileContents)
	if err != nil {
		return fmt.Errorf("Failed to download GPU driver build %s, %v", build.DriverPath, err)
	}
	fileName := cloudcommon.GetGPUDriverBuildStoragePath(driverKey, build.Name, ext)

	// If Linux, then validate the pkg
	//     * Pkg must be deb pkg
	//     * Pkg control file must have kernel dependency specified
	if build.OperatingSystem == edgeproto.OSType_LINUX {
		cb.Send(&edgeproto.Result{Message: "Verify if GPU driver package is a debian package"})
		if ext != ".deb" {
			return fmt.Errorf("Only supported file extension for Linux GPU driver is '.deb', given %s", ext)
		}
		cb.Send(&edgeproto.Result{Message: "Verify if kernel dependency is specified as part of package's control file"})
		testFileName := "/tmp/" + strings.ReplaceAll(fileName, "/", "_")
		err = ioutil.WriteFile(testFileName, []byte(fileContents), 0644)
		if err != nil {
			return fmt.Errorf("Failed to write driver package to %s, %v", testFileName, err)
		}
		defer os.Remove(testFileName)
		localClient := &pc.LocalClient{}
		cmd := fmt.Sprintf("dpkg-deb -I %s | grep -i 'Depends: linux-image-%s'", testFileName, build.KernelVersion)
		out, err := localClient.Output(cmd)
		if err != nil && out != "" {
			return fmt.Errorf("Invalid driver package(%q), should be a valid debian package, %s, %v", fileName, out, err)
		}
		if out == "" {
			return fmt.Errorf("Driver package(%q) should have Linux Kernel dependency(%q) specified as part of debian control file, %v", fileName, build.KernelVersion, err)
		}
	}

	// Upload the package to GCS
	cb.Send(&edgeproto.Result{Message: "Upload the GPU driver to secure storage"})
	err = storageClient.UploadObject(ctx, fileName, bytes.NewBufferString(fileContents))
	if err != nil {
		return fmt.Errorf("Failed to upload GPU driver build to %s, %v", fileName, err)
	}
	build.DriverPath = cloudcommon.GetGPUDriverURL(driverKey, nodeMgr.DeploymentTag, build.Name, ext)
	return nil
}

func setupGPUDriverLicenseConfig(ctx context.Context, storageClient *gcs.GCSClient, driverKey *edgeproto.GPUDriverKey, licenseConfig *string, cb edgeproto.GPUDriverApi_CreateGPUDriverServer) error {
	cb.Send(&edgeproto.Result{Message: "Upload the GPU driver license config to secure storage"})
	fileName := cloudcommon.GetGPUDriverLicenseStoragePath(driverKey)
	err := storageClient.UploadObject(ctx, fileName, bytes.NewBufferString(*licenseConfig))
	if err != nil {
		return fmt.Errorf("Failed to upload GPU driver license to %s, %v", fileName, err)
	}
	*licenseConfig = cloudcommon.GetGPUDriverLicenseURL(driverKey, nodeMgr.DeploymentTag)
	return nil
}

func deleteGPUDriverLicenseConfig(ctx context.Context, storageClient *gcs.GCSClient, driverKey *edgeproto.GPUDriverKey) error {
	fileName := cloudcommon.GetGPUDriverLicenseStoragePath(driverKey)
	err := storageClient.DeleteObject(ctx, fileName)
	if err != nil {
		return fmt.Errorf("Failed to delete GPU driver license to %s, %v", fileName, err)
	}
	return nil
}

func (s *GPUDriverApi) CreateGPUDriver(in *edgeproto.GPUDriver, cb edgeproto.GPUDriverApi_CreateGPUDriverServer) (reterr error) {
	ctx := cb.Context()
	if err := in.Validate(nil); err != nil {
		return err
	}

	if in.Key.Organization == "" {
		// Public GPU drivers have no org associated with them
	}

	// To ensure updates to etcd and GCS happens atomically:
	// Step-1: First commit to etcd
	// Step-2: Validate and upload the builds/license-config to GCS
	// Step-3: And then update build details to reflect GCS URL and update it to etcd

	// Step-1: First commit to etcd
	err := s.sync.ApplySTMWait(ctx, func(stm concurrency.STM) error {
		if s.store.STMGet(stm, &in.Key, nil) {
			return in.Key.ExistsError()
		}
		in.State = ChangeInProgress
		s.store.STMPut(stm, in)
		return nil
	})
	if err != nil {
		return err
	}

	defer func() {
		if reterr != nil {
			// undo changes
			err = s.DeleteGPUDriver(in, cb)
			if err != nil {
				log.SpanLog(ctx, log.DebugLevelApi, "failed to undo gpu driver create", "key", in.Key, "err", err)
			}
		}
	}()

	if len(in.Builds) > 0 || in.LicenseConfig != "" {
		storageClient, err := getGCSStorageClient(ctx)
		if err != nil {
			return err
		}
		defer storageClient.Close()

		// Step-2: Validate and upload the builds/license-config to GCS
		for ii, build := range in.Builds {
			cb.Send(&edgeproto.Result{Message: "Setting up GPU driver build " + build.Name})
			err := setupGPUDriver(ctx, storageClient, &in.Key, &build, cb)
			if err != nil {
				return err
			}
			in.Builds[ii].DriverPath = build.DriverPath
		}

		// If license config is present, upload it to GCS
		if in.LicenseConfig != "" {
			err := setupGPUDriverLicenseConfig(ctx, storageClient, &in.Key, &in.LicenseConfig, cb)
			if err != nil {
				return err
			}
		}
	}

	// Step-3: And then update build details to reflect GCS URL and update it to etcd
	in.State = ""
	_, err = s.store.Put(ctx, in, s.sync.syncWait)
	if err != nil {
		return err
	}
	cb.Send(&edgeproto.Result{Message: "GPU driver created successfully"})
	return nil
}

func (s *GPUDriverApi) UpdateGPUDriver(in *edgeproto.GPUDriver, cb edgeproto.GPUDriverApi_UpdateGPUDriverServer) error {
	ctx := cb.Context()
	err := in.ValidateUpdateFields()
	if err != nil {
		return err
	}
	fmap := edgeproto.MakeFieldMap(in.Fields)
	if err := in.Validate(fmap); err != nil {
		return err
	}

	ignoreState := in.IgnoreState
	in.IgnoreState = false

	// To ensure updates to etcd and GCS happens atomically:
	// Step-1: First commit to etcd
	// Step-2: Validate and upload the license-config to GCS
	// Step-3: And then update build details to reflect GCS URL and update it to etcd

	// Step-1: First commit to etcd
	changed := 0
	err = s.sync.ApplySTMWait(ctx, func(stm concurrency.STM) error {
		cur := edgeproto.GPUDriver{}
		changed = 0
		if !s.store.STMGet(stm, &in.Key, &cur) {
			return in.Key.NotFoundError()
		}
		if err := isBusyState(&in.Key, cur.State, ignoreState); err != nil {
			return err
		}
		changed = cur.CopyInFields(in)
		if err := cur.Validate(nil); err != nil {
			return err
		}
		if changed == 0 {
			return nil
		}
		cur.State = ChangeInProgress
		s.store.STMPut(stm, &cur)
		return nil
	})
	if err != nil {
		return err
	}
	if changed == 0 {
		return nil
	}

	// Step-2: Validate and upload the license-config to GCS
	// If license config is present, upload it to GCS
	if _, found := fmap[edgeproto.GPUDriverFieldLicenseConfig]; found {
		storageClient, err := getGCSStorageClient(ctx)
		if err != nil {
			return err
		}
		defer storageClient.Close()
		if in.LicenseConfig == "" {
			cb.Send(&edgeproto.Result{Message: "Deleting GPU driver license config from secure storage"})
			// Delete license config from GCS
			err = deleteGPUDriverLicenseConfig(ctx, storageClient, &in.Key)
			if err != nil {
				return err
			}
		} else {
			err = setupGPUDriverLicenseConfig(ctx, storageClient, &in.Key, &in.LicenseConfig, cb)
			if err != nil {
				return err
			}
		}
	}

	// Step-3: And then update license config to reflect GCS URL and update it to etcd
	err = s.sync.ApplySTMWait(ctx, func(stm concurrency.STM) error {
		cur := edgeproto.GPUDriver{}
		changed := 0
		if !s.store.STMGet(stm, &in.Key, &cur) {
			return in.Key.NotFoundError()
		}
		changed = cur.CopyInFields(in)
		if changed == 0 {
			return nil
		}
		cur.State = ""
		s.store.STMPut(stm, &cur)
		return nil
	})
	if err != nil {
		return err
	}
	cb.Send(&edgeproto.Result{Message: "GPU driver updated successfully"})
	return nil
}

func isBusyState(key *edgeproto.GPUDriverKey, state string, ignoreState bool) error {
	if !ignoreState && state == ChangeInProgress {
		return fmt.Errorf("An action is already in progress for GPU driver %s", key.String())
	}
	return nil
}

func (s *GPUDriverApi) DeleteGPUDriver(in *edgeproto.GPUDriver, cb edgeproto.GPUDriverApi_DeleteGPUDriverServer) error {
	ctx := cb.Context()
	if err := in.Key.ValidateKey(); err != nil {
		return err
	}
	// Validate if driver is in use by Cloudlet
	if cloudletApi.UsesGPUDriver(&in.Key) {
		return fmt.Errorf("GPU driver in use by Cloudlet")
	}

	ignoreState := in.IgnoreState
	in.IgnoreState = false

	// To ensure updates to etcd and GCS happens atomically:
	// Step-1: First update state in etcd
	// Step-2: Delete objects from GCS
	// Step-3: And then delete obj from etcd

	// Step-1: First update state in etcd
	buildFiles := []string{}
	licenseConfig := ""
	err := s.sync.ApplySTMWait(ctx, func(stm concurrency.STM) error {
		cur := edgeproto.GPUDriver{}
		if !s.store.STMGet(stm, &in.Key, &cur) {
			return in.Key.NotFoundError()
		}
		if err := isBusyState(&in.Key, cur.State, ignoreState); err != nil {
			return err
		}
		buildFiles = []string{}
		for _, build := range in.Builds {
			fileName := cloudcommon.GetGPUDriverBuildPathFromURL(build.DriverPath, nodeMgr.DeploymentTag)
			buildFiles = append(buildFiles, fileName)
		}
		licenseConfig = cur.LicenseConfig
		cur.State = ChangeInProgress
		s.store.STMPut(stm, &cur)
		return nil
	})
	if err != nil {
		return err
	}
	// Step-2: Delete objects from GCS
	if len(buildFiles) > 0 || licenseConfig != "" {
		storageClient, err := getGCSStorageClient(ctx)
		if err != nil {
			return err
		}
		defer storageClient.Close()
		// Delete license config from GCS
		err = deleteGPUDriverLicenseConfig(ctx, storageClient, &in.Key)
		if err != nil {
			return err
		}
		// Delete builds from GCS
		for _, filename := range buildFiles {
			err = storageClient.DeleteObject(ctx, filename)
			if err != nil {
				return fmt.Errorf("Failed to delete GPU driver build %s, %v: %v", filename, in.Key, err)
			}
		}
	}

	// Step-3: And then delete obj from etcd
	_, err = s.store.Delete(ctx, in, s.sync.syncWait)
	if err != nil {
		return err
	}
	cb.Send(&edgeproto.Result{Message: "GPU driver deleted successfully"})
	return nil
}

func (s *GPUDriverApi) ShowGPUDriver(in *edgeproto.GPUDriver, cb edgeproto.GPUDriverApi_ShowGPUDriverServer) error {
	return s.cache.Show(in, func(obj *edgeproto.GPUDriver) error {
		err := cb.Send(obj)
		return err
	})
}

func (s *GPUDriverApi) AddGPUDriverBuild(in *edgeproto.GPUDriverBuildMember, cb edgeproto.GPUDriverApi_AddGPUDriverBuildServer) (reterr error) {
	ctx := cb.Context()
	if err := in.Validate(); err != nil {
		return err
	}

	ignoreState := in.IgnoreState
	in.IgnoreState = false

	// To ensure updates to etcd and GCS happens atomically:
	// Step-1: First commit to etcd
	// Step-2: Validate and upload the build to GCS
	// Step-3: And then update build details to reflect GCS URL and update it to etcd

	// Step-1: First commit to etcd
	err := s.sync.ApplySTMWait(ctx, func(stm concurrency.STM) error {
		cur := edgeproto.GPUDriver{}
		if !s.store.STMGet(stm, &in.Key, &cur) {
			return in.Key.NotFoundError()
		}
		if err := isBusyState(&in.Key, cur.State, ignoreState); err != nil {
			return err
		}
		for ii, _ := range cur.Builds {
			if cur.Builds[ii].Name == in.Build.Name {
				return fmt.Errorf("GPU driver build with same name already exists")
			}
		}
		cur.Builds = append(cur.Builds, in.Build)
		cur.State = ChangeInProgress
		s.store.STMPut(stm, &cur)
		return nil
	})
	if err != nil {
		return err
	}

	defer func() {
		if reterr != nil {
			// undo changes
			err = s.RemoveGPUDriverBuild(in, cb)
			if err != nil {
				log.SpanLog(ctx, log.DebugLevelApi, "failed to undo gpu driver build", "key", in.Key, "err", err)
			}
		}
	}()

	// Step-2: Validate and upload the build to GCS
	storageClient, err := getGCSStorageClient(ctx)
	if err != nil {
		return err
	}
	defer storageClient.Close()

	err = setupGPUDriver(ctx, storageClient, &in.Key, &in.Build, cb)
	if err != nil {
		return err
	}

	// Step-3: And then update build details to reflect GCS URL and update it to etcd
	err = s.sync.ApplySTMWait(ctx, func(stm concurrency.STM) error {
		cur := edgeproto.GPUDriver{}
		if !s.store.STMGet(stm, &in.Key, &cur) {
			return in.Key.NotFoundError()
		}
		found := false
		for ii, _ := range cur.Builds {
			if cur.Builds[ii].Name == in.Build.Name {
				cur.Builds[ii] = in.Build
				found = true
				break
			}
		}
		if !found {
			cur.Builds = append(cur.Builds, in.Build)
		}
		cur.State = ""
		s.store.STMPut(stm, &cur)
		return nil
	})
	if err != nil {
		return err
	}
	cb.Send(&edgeproto.Result{Message: "GPU driver build added successfully"})
	return nil
}

func (s *GPUDriverApi) RemoveGPUDriverBuild(in *edgeproto.GPUDriverBuildMember, cb edgeproto.GPUDriverApi_RemoveGPUDriverBuildServer) error {
	ctx := cb.Context()
	if err := in.Key.ValidateKey(); err != nil {
		return err
	}
	if err := in.Build.ValidateName(); err != nil {
		return err
	}
	ignoreState := in.IgnoreState
	in.IgnoreState = false

	// To ensure updates to etcd and GCS happens atomically:
	// Step-1: First commit to etcd
	// Step-2: Delete build from GCS
	// Step-3: And then update build details in etcd to reflect build deletion

	// Step-1: First commit to etcd
	driverURL := ""
	err := s.sync.ApplySTMWait(ctx, func(stm concurrency.STM) error {
		cur := edgeproto.GPUDriver{}
		if !s.store.STMGet(stm, &in.Key, &cur) {
			return in.Key.NotFoundError()
		}
		if err := isBusyState(&in.Key, cur.State, ignoreState); err != nil {
			return err
		}
		changed := false
		for ii, build := range cur.Builds {
			if build.Name == in.Build.Name {
				cur.Builds = append(cur.Builds[:ii], cur.Builds[ii+1:]...)
				driverURL = build.DriverPath
				changed = true
				break
			}
		}
		if !changed {
			return nil
		}
		cur.State = ChangeInProgress
		s.store.STMPut(stm, &cur)
		return nil
	})
	if err != nil {
		return err
	}
	if driverURL == "" {
		return nil
	}
	// Step-2: Delete build from GCS
	cb.Send(&edgeproto.Result{Message: "Deleting GPU driver build from secure storage"})
	fileName := cloudcommon.GetGPUDriverBuildPathFromURL(driverURL, nodeMgr.DeploymentTag)
	storageClient, err := getGCSStorageClient(ctx)
	if err != nil {
		return err
	}
	defer storageClient.Close()
	err = storageClient.DeleteObject(ctx, fileName)
	if err != nil {
		log.SpanLog(ctx, log.DebugLevelApi, "failed to delete GCS object", "filename", fileName, "err", err)
		return err
	}
	// Step-3: And then update build details in etcd to reflect build deletion
	err = s.sync.ApplySTMWait(ctx, func(stm concurrency.STM) error {
		cur := edgeproto.GPUDriver{}
		if !s.store.STMGet(stm, &in.Key, &cur) {
			return in.Key.NotFoundError()
		}
		cur.State = ""
		s.store.STMPut(stm, &cur)
		return nil
	})
	if err != nil {
		return err
	}
	cb.Send(&edgeproto.Result{Message: "GPU driver build removed successfully"})
	return nil
}

func (s *GPUDriverApi) GetGPUDriverBuildURL(ctx context.Context, in *edgeproto.GPUDriverBuildMember) (*edgeproto.GPUDriverBuildURL, error) {
	if err := in.Key.ValidateKey(); err != nil {
		return &edgeproto.GPUDriverBuildURL{}, err
	}
	if err := in.Build.ValidateName(); err != nil {
		return &edgeproto.GPUDriverBuildURL{}, err
	}
	var driverURL string
	err := s.sync.ApplySTMWait(ctx, func(stm concurrency.STM) error {
		cur := edgeproto.GPUDriver{}
		if !s.store.STMGet(stm, &in.Key, &cur) {
			return in.Key.NotFoundError()
		}
		ignoreState := true
		if err := isBusyState(&in.Key, cur.State, !ignoreState); err != nil {
			return err
		}
		for _, build := range cur.Builds {
			if build.Name == in.Build.Name {
				driverURL = build.DriverPath
				break
			}
		}
		return nil
	})
	if err != nil {
		return &edgeproto.GPUDriverBuildURL{}, err
	}
	if driverURL == "" {
		return &edgeproto.GPUDriverBuildURL{}, fmt.Errorf("GPUDriver build %s for driver %s not found", in.Build.Name, in.Key.GetKeyString())
	}
	fileName := cloudcommon.GetGPUDriverBuildPathFromURL(driverURL, nodeMgr.DeploymentTag)
	storageClient, err := getGCSStorageClient(ctx)
	if err != nil {
		return nil, err
	}
	defer storageClient.Close()
	signedUrl, err := storageClient.GenerateV4GetObjectSignedURL(ctx, fileName, GPUDriverBuildURLValidity)
	if err != nil {
		log.SpanLog(ctx, log.DebugLevelApi, "failed to generate signed URL for GCS object", "filename", fileName, "err", err)
		return nil, err
	}
	return &edgeproto.GPUDriverBuildURL{
		BuildUrlPath: signedUrl,
		Validity:     edgeproto.Duration(GPUDriverBuildURLValidity),
	}, nil
}