package pfutils

import (
	"fmt"
	"os"
	"plugin"

	pf "github.com/mobiledgex/edge-cloud/cloud-resource-manager/platform"
	"github.com/mobiledgex/edge-cloud/cloud-resource-manager/platform/dind"
	"github.com/mobiledgex/edge-cloud/cloud-resource-manager/platform/fake"
	"github.com/mobiledgex/edge-cloud/log"
)

var solib = ""

func GetPlatform(plat string) (pf.Platform, error) {
	// Building plugins is slow, so directly importable
	// platforms are not built as plugins.
	if plat == "PLATFORM_TYPE_DIND" {
		return &dind.Platform{}, nil
	} else if plat == "PLATFORM_TYPE_FAKE" {
		return &fake.Platform{}, nil
	}

	// Load platform from plugin
	if solib == "" {
		solib = os.Getenv("GOPATH") + "/plugins/platforms.so"
	}
	log.DebugLog(log.DebugLevelMexos, "Loading plugin", "plugin", solib)
	plug, err := plugin.Open(solib)
	if err != nil {
		log.DebugLog(log.DebugLevelMexos, "failed to load plugin", "plugin", solib, "platform", plat, "error", err)
		return nil, fmt.Errorf("failed to load plugin for platform: %s, err: %v", plat, err)
	}
	sym, err := plug.Lookup("GetPlatform")
	if err != nil {
		log.DebugLog(log.DebugLevelMexos, "plugin does not have GetPlatform symbol", "plugin", solib)
		return nil, fmt.Errorf("failed to load plugin for platform: %s, err: GetPlatform symbol not found", plat)
	}
	getPlatFunc, ok := sym.(func(plat string) (pf.Platform, error))
	if !ok {
		log.DebugLog(log.DebugLevelMexos, "plugin GetPlatform symbol does not implement func(plat string) (platform.Platform, error)", "plugin", solib)
		return nil, fmt.Errorf("failed to load plugin for platform: %s, err: GetPlatform symbol does not implement func(plat string) (platform.Platform, error)", plat)
	}
	log.DebugLog(log.DebugLevelMexos, "Creating platform")
	return getPlatFunc(plat)
}