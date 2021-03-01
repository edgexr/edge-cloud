package kind

import (
	"context"

	"github.com/mobiledgex/edge-cloud/cloud-resource-manager/platform"
	"github.com/mobiledgex/edge-cloud/cloud-resource-manager/platform/common/xind"
	"github.com/mobiledgex/edge-cloud/edgeproto"
)

type Platform struct {
	xind.Xind
}

func (s *Platform) Init(ctx context.Context, platformConfig *platform.PlatformConfig, caches *platform.Caches, updateCallback edgeproto.CacheUpdateCallback) error {
	return s.Xind.Init(ctx, platformConfig, caches, s, updateCallback)
}