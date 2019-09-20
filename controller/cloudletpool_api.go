package main

import (
	"context"
	"fmt"

	"github.com/coreos/etcd/clientv3/concurrency"
	"github.com/mobiledgex/edge-cloud/cloudcommon"
	"github.com/mobiledgex/edge-cloud/edgeproto"
	"github.com/mobiledgex/edge-cloud/objstore"
)

type CloudletPoolApi struct {
	sync  *Sync
	store edgeproto.CloudletPoolStore
	cache edgeproto.CloudletPoolCache
}

var cloudletPoolApi = CloudletPoolApi{}

func InitCloudletPoolApi(sync *Sync) {
	cloudletPoolApi.sync = sync
	cloudletPoolApi.store = edgeproto.NewCloudletPoolStore(sync.store)
	edgeproto.InitCloudletPoolCache(&cloudletPoolApi.cache)
	sync.RegisterCache(&cloudletPoolApi.cache)
}

func (s *CloudletPoolApi) registerPublicPool(ctx context.Context) error {
	err := s.sync.ApplySTMWait(ctx, func(stm concurrency.STM) error {
		pool := edgeproto.CloudletPool{}
		pool.Key.Name = cloudcommon.PublicCloudletPool
		if s.store.STMGet(stm, &pool.Key, &pool) {
			// already present
			return nil
		}
		s.store.STMPut(stm, &pool)
		return nil
	})
	return err
}

func (s *CloudletPoolApi) CreateCloudletPool(ctx context.Context, in *edgeproto.CloudletPool) (*edgeproto.Result, error) {
	if err := in.Validate(edgeproto.CloudletPoolAllFieldsMap); err != nil {
		return &edgeproto.Result{}, err
	}

	err := s.sync.ApplySTMWait(ctx, func(stm concurrency.STM) error {
		if s.store.STMGet(stm, &in.Key, nil) {
			return objstore.ErrKVStoreKeyExists
		}
		s.store.STMPut(stm, in)
		return nil
	})
	return &edgeproto.Result{}, err
}

func (s *CloudletPoolApi) DeleteCloudletPool(ctx context.Context, in *edgeproto.CloudletPool) (*edgeproto.Result, error) {
	if in.Key.Name == cloudcommon.PublicCloudletPool {
		return &edgeproto.Result{}, fmt.Errorf("cannot delete Public pool")
	}
	err := s.sync.ApplySTMWait(ctx, func(stm concurrency.STM) error {
		if !s.store.STMGet(stm, &in.Key, nil) {
			return objstore.ErrKVStoreKeyNotFound
		}
		s.store.STMDel(stm, &in.Key)
		return nil
	})
	if err == nil {
		cloudletPoolMemberApi.poolDeleted(ctx, &in.Key)
	}
	return &edgeproto.Result{}, err
}

func (s *CloudletPoolApi) ShowCloudletPool(in *edgeproto.CloudletPool, cb edgeproto.CloudletPoolApi_ShowCloudletPoolServer) error {
	err := s.cache.Show(in, func(obj *edgeproto.CloudletPool) error {
		err := cb.Send(obj)
		return err
	})
	return err
}

func (s *CloudletPoolApi) showPoolsByKeys(keys map[edgeproto.CloudletPoolKey]struct{}, cb func(obj *edgeproto.CloudletPool) error) error {
	s.cache.Mux.Lock()
	defer s.cache.Mux.Unlock()

	for key, obj := range s.cache.Objs {
		if _, found := keys[key]; !found {
			continue
		}
		err := cb(obj)
		if err != nil {
			return err
		}
	}
	return nil
}