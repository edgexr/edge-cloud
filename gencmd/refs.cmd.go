// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: refs.proto

package gencmd

import (
	"context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	"github.com/mobiledgex/edge-cloud/cli"
	edgeproto "github.com/mobiledgex/edge-cloud/edgeproto"
	_ "github.com/mobiledgex/edge-cloud/protogen"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/status"
	"io"
	math "math"
	"strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Auto-generated code: DO NOT EDIT
var CloudletRefsApiCmd edgeproto.CloudletRefsApiClient

var ShowCloudletRefsCmd = &cli.Command{
	Use:          "ShowCloudletRefs",
	OptionalArgs: strings.Join(append(CloudletRefsRequiredArgs, CloudletRefsOptionalArgs...), " "),
	AliasArgs:    strings.Join(CloudletRefsAliasArgs, " "),
	SpecialArgs:  &CloudletRefsSpecialArgs,
	Comments:     CloudletRefsComments,
	ReqData:      &edgeproto.CloudletRefs{},
	ReplyData:    &edgeproto.CloudletRefs{},
	Run:          runShowCloudletRefs,
}

func runShowCloudletRefs(c *cli.Command, args []string) error {
	if cli.SilenceUsage {
		c.CobraCmd.SilenceUsage = true
	}
	obj := c.ReqData.(*edgeproto.CloudletRefs)
	_, err := c.ParseInput(args)
	if err != nil {
		return err
	}
	return ShowCloudletRefs(c, obj)
}

func ShowCloudletRefs(c *cli.Command, in *edgeproto.CloudletRefs) error {
	if CloudletRefsApiCmd == nil {
		return fmt.Errorf("CloudletRefsApi client not initialized")
	}
	ctx := context.Background()
	stream, err := CloudletRefsApiCmd.ShowCloudletRefs(ctx, in)
	if err != nil {
		errstr := err.Error()
		st, ok := status.FromError(err)
		if ok {
			errstr = st.Message()
		}
		return fmt.Errorf("ShowCloudletRefs failed: %s", errstr)
	}

	objs := make([]*edgeproto.CloudletRefs, 0)
	for {
		obj, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			errstr := err.Error()
			st, ok := status.FromError(err)
			if ok {
				errstr = st.Message()
			}
			return fmt.Errorf("ShowCloudletRefs recv failed: %s", errstr)
		}
		objs = append(objs, obj)
	}
	if len(objs) == 0 {
		return nil
	}
	c.WriteOutput(c.CobraCmd.OutOrStdout(), objs, cli.OutputFormat)
	return nil
}

// this supports "Create" and "Delete" commands on ApplicationData
func ShowCloudletRefss(c *cli.Command, data []edgeproto.CloudletRefs, err *error) {
	if *err != nil {
		return
	}
	for ii, _ := range data {
		fmt.Printf("ShowCloudletRefs %v\n", data[ii])
		myerr := ShowCloudletRefs(c, &data[ii])
		if myerr != nil {
			*err = myerr
			break
		}
	}
}

var CloudletRefsApiCmds = []*cobra.Command{
	ShowCloudletRefsCmd.GenCmd(),
}

var ClusterRefsApiCmd edgeproto.ClusterRefsApiClient

var ShowClusterRefsCmd = &cli.Command{
	Use:          "ShowClusterRefs",
	OptionalArgs: strings.Join(append(ClusterRefsRequiredArgs, ClusterRefsOptionalArgs...), " "),
	AliasArgs:    strings.Join(ClusterRefsAliasArgs, " "),
	SpecialArgs:  &ClusterRefsSpecialArgs,
	Comments:     ClusterRefsComments,
	ReqData:      &edgeproto.ClusterRefs{},
	ReplyData:    &edgeproto.ClusterRefs{},
	Run:          runShowClusterRefs,
}

func runShowClusterRefs(c *cli.Command, args []string) error {
	if cli.SilenceUsage {
		c.CobraCmd.SilenceUsage = true
	}
	obj := c.ReqData.(*edgeproto.ClusterRefs)
	_, err := c.ParseInput(args)
	if err != nil {
		return err
	}
	return ShowClusterRefs(c, obj)
}

func ShowClusterRefs(c *cli.Command, in *edgeproto.ClusterRefs) error {
	if ClusterRefsApiCmd == nil {
		return fmt.Errorf("ClusterRefsApi client not initialized")
	}
	ctx := context.Background()
	stream, err := ClusterRefsApiCmd.ShowClusterRefs(ctx, in)
	if err != nil {
		errstr := err.Error()
		st, ok := status.FromError(err)
		if ok {
			errstr = st.Message()
		}
		return fmt.Errorf("ShowClusterRefs failed: %s", errstr)
	}

	objs := make([]*edgeproto.ClusterRefs, 0)
	for {
		obj, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			errstr := err.Error()
			st, ok := status.FromError(err)
			if ok {
				errstr = st.Message()
			}
			return fmt.Errorf("ShowClusterRefs recv failed: %s", errstr)
		}
		objs = append(objs, obj)
	}
	if len(objs) == 0 {
		return nil
	}
	c.WriteOutput(c.CobraCmd.OutOrStdout(), objs, cli.OutputFormat)
	return nil
}

// this supports "Create" and "Delete" commands on ApplicationData
func ShowClusterRefss(c *cli.Command, data []edgeproto.ClusterRefs, err *error) {
	if *err != nil {
		return
	}
	for ii, _ := range data {
		fmt.Printf("ShowClusterRefs %v\n", data[ii])
		myerr := ShowClusterRefs(c, &data[ii])
		if myerr != nil {
			*err = myerr
			break
		}
	}
}

var ClusterRefsApiCmds = []*cobra.Command{
	ShowClusterRefsCmd.GenCmd(),
}

var AppInstRefsApiCmd edgeproto.AppInstRefsApiClient

var ShowAppInstRefsCmd = &cli.Command{
	Use:          "ShowAppInstRefs",
	OptionalArgs: strings.Join(append(AppInstRefsRequiredArgs, AppInstRefsOptionalArgs...), " "),
	AliasArgs:    strings.Join(AppInstRefsAliasArgs, " "),
	SpecialArgs:  &AppInstRefsSpecialArgs,
	Comments:     AppInstRefsComments,
	ReqData:      &edgeproto.AppInstRefs{},
	ReplyData:    &edgeproto.AppInstRefs{},
	Run:          runShowAppInstRefs,
}

func runShowAppInstRefs(c *cli.Command, args []string) error {
	if cli.SilenceUsage {
		c.CobraCmd.SilenceUsage = true
	}
	obj := c.ReqData.(*edgeproto.AppInstRefs)
	_, err := c.ParseInput(args)
	if err != nil {
		return err
	}
	return ShowAppInstRefs(c, obj)
}

func ShowAppInstRefs(c *cli.Command, in *edgeproto.AppInstRefs) error {
	if AppInstRefsApiCmd == nil {
		return fmt.Errorf("AppInstRefsApi client not initialized")
	}
	ctx := context.Background()
	stream, err := AppInstRefsApiCmd.ShowAppInstRefs(ctx, in)
	if err != nil {
		errstr := err.Error()
		st, ok := status.FromError(err)
		if ok {
			errstr = st.Message()
		}
		return fmt.Errorf("ShowAppInstRefs failed: %s", errstr)
	}

	objs := make([]*edgeproto.AppInstRefs, 0)
	for {
		obj, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			errstr := err.Error()
			st, ok := status.FromError(err)
			if ok {
				errstr = st.Message()
			}
			return fmt.Errorf("ShowAppInstRefs recv failed: %s", errstr)
		}
		objs = append(objs, obj)
	}
	if len(objs) == 0 {
		return nil
	}
	c.WriteOutput(c.CobraCmd.OutOrStdout(), objs, cli.OutputFormat)
	return nil
}

// this supports "Create" and "Delete" commands on ApplicationData
func ShowAppInstRefss(c *cli.Command, data []edgeproto.AppInstRefs, err *error) {
	if *err != nil {
		return
	}
	for ii, _ := range data {
		fmt.Printf("ShowAppInstRefs %v\n", data[ii])
		myerr := ShowAppInstRefs(c, &data[ii])
		if myerr != nil {
			*err = myerr
			break
		}
	}
}

var AppInstRefsApiCmds = []*cobra.Command{
	ShowAppInstRefsCmd.GenCmd(),
}

var VMResourceRequiredArgs = []string{
	"key.clusterkey.name",
	"key.cloudletkey.organization",
	"key.cloudletkey.name",
	"key.organization",
}
var VMResourceOptionalArgs = []string{
	"vmflavor.name",
	"vmflavor.vcpus",
	"vmflavor.ram",
	"vmflavor.disk",
	"vmflavor.propmap",
	"type",
	"appaccesstype",
}
var VMResourceAliasArgs = []string{}
var VMResourceComments = map[string]string{
	"key.clusterkey.name":          "Cluster name",
	"key.cloudletkey.organization": "Organization of the cloudlet site",
	"key.cloudletkey.name":         "Name of the cloudlet",
	"key.organization":             "Name of Developer organization that this cluster belongs to",
	"vmflavor.name":                "Name of the flavor on the Cloudlet",
	"vmflavor.vcpus":               "Number of VCPU cores on the Cloudlet",
	"vmflavor.ram":                 "Ram in MB on the Cloudlet",
	"vmflavor.disk":                "Amount of disk in GB on the Cloudlet",
	"vmflavor.propmap":             "OS Flavor Properties, if any",
	"type":                         "Resource Type can be platform, rootlb, cluster-master, cluster-k8s-node, cluster-docker-node, appvm",
	"appaccesstype":                "Access type for resource of type App VM, one of DefaultForDeployment, Direct, LoadBalancer",
}
var VMResourceSpecialArgs = map[string]string{
	"vmflavor.propmap": "StringToString",
}
var CloudletRefsRequiredArgs = []string{
	"key.organization",
	"key.name",
}
var CloudletRefsOptionalArgs = []string{
	"rootlbports:#.key",
	"rootlbports:#.value",
	"useddynamicips",
	"usedstaticips",
	"optresusedmap:#.key",
	"optresusedmap:#.value",
	"reservedautoclusterids",
	"clusterinsts:#.clusterkey.name",
	"clusterinsts:#.organization",
	"vmappinsts:#.appkey.organization",
	"vmappinsts:#.appkey.name",
	"vmappinsts:#.appkey.version",
	"vmappinsts:#.clusterinstkey.clusterkey.name",
	"vmappinsts:#.clusterinstkey.organization",
}
var CloudletRefsAliasArgs = []string{}
var CloudletRefsComments = map[string]string{
	"key.organization":                            "Organization of the cloudlet site",
	"key.name":                                    "Name of the cloudlet",
	"useddynamicips":                              "Used dynamic IPs",
	"usedstaticips":                               "Used static IPs",
	"reservedautoclusterids":                      "Track reservable autoclusterinsts ids in use. This is a bitmap.",
	"clusterinsts:#.clusterkey.name":              "Cluster name",
	"clusterinsts:#.organization":                 "Name of Developer organization that this cluster belongs to",
	"vmappinsts:#.appkey.organization":            "App developer organization",
	"vmappinsts:#.appkey.name":                    "App name",
	"vmappinsts:#.appkey.version":                 "App version",
	"vmappinsts:#.clusterinstkey.clusterkey.name": "Cluster name",
	"vmappinsts:#.clusterinstkey.organization":    "Name of Developer organization that this cluster belongs to",
}
var CloudletRefsSpecialArgs = map[string]string{}
var ClusterRefsRequiredArgs = []string{
	"key.clusterkey.name",
	"key.cloudletkey.organization",
	"key.cloudletkey.name",
	"key.organization",
}
var ClusterRefsOptionalArgs = []string{
	"apps:#.organization",
	"apps:#.name",
	"apps:#.version",
	"usedram",
	"usedvcores",
	"useddisk",
}
var ClusterRefsAliasArgs = []string{}
var ClusterRefsComments = map[string]string{
	"key.clusterkey.name":          "Cluster name",
	"key.cloudletkey.organization": "Organization of the cloudlet site",
	"key.cloudletkey.name":         "Name of the cloudlet",
	"key.organization":             "Name of Developer organization that this cluster belongs to",
	"apps:#.organization":          "App developer organization",
	"apps:#.name":                  "App name",
	"apps:#.version":               "App version",
	"usedram":                      "Used RAM in MB",
	"usedvcores":                   "Used VCPU cores",
	"useddisk":                     "Used disk in GB",
}
var ClusterRefsSpecialArgs = map[string]string{}
var AppInstRefsRequiredArgs = []string{
	"key.organization",
	"key.name",
	"key.version",
}
var AppInstRefsOptionalArgs = []string{
	"insts:#.key",
	"insts:#.value",
}
var AppInstRefsAliasArgs = []string{}
var AppInstRefsComments = map[string]string{
	"key.organization": "App developer organization",
	"key.name":         "App name",
	"key.version":      "App version",
}
var AppInstRefsSpecialArgs = map[string]string{}
