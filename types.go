package raftinspector

import (
	"time"
)

type NodeInfo struct {
	ID                string
	Address           string
	State             string
	Term              string
	CommitIndex       string
	AppliedIndex      string
	LastSnapshotIndex string
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
