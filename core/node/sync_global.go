package node

import (
	"github.com/snple/kokomi/pb/nodes"
)

type SyncGlobalService struct {
	ns *NodeService

	nodes.UnimplementedSyncGlobalServiceServer
}

func newSyncGlobalService(ns *NodeService) *SyncGlobalService {
	return &SyncGlobalService{
		ns: ns,
	}
}
