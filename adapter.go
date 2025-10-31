package raftinspector

import "context"

type Adapter interface {
	GetNodeInfo(ctx context.Context) NodeInfo
	GetClusterInfo(ctx context.Context) ClusterInfo
}
