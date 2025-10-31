package raftinspector

import "context"

type Adapter interface {
	NodeInfo(ctx context.Context) NodeInfo
	ClusterInfo(ctx context.Context) ClusterInfo
}
