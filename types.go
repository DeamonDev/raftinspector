package raftinspector

import (
	"time"
)

type NodeInfo struct {
	ID                string
	Address           string
	State             string
	Term              uint64
	CommitIndex       uint64
	AppliedIndex      uint64
	LastSnapshotIndex uint64
}

type ClusterInfo struct {
	NodesID  []string
	LeaderID string
	Term     uint64
}

type Metrics struct {
	Timestamp    time.Time
	LeaderID     string
	AppliedIndex uint64
}
