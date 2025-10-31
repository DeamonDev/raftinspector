package hashicorp

import (
	"context"
	"log"
	"strconv"

	"github.com/deamondev/raftinspector"
	"github.com/hashicorp/raft"
)

/*
	HashiCorpAdapter Use this adapter if you are able to instrument your raft nodes.

If you're not able to modify code, then you should use a metrics server adapter
In that case your adapter will be readonly.
This adapter allows modifying the state of the cluster.
*/
type HashiCorpAdapter struct {
	Node    *raft.Raft
	Address string
}

func NewHashiCorpAdapter(r *raft.Raft, a string) *HashiCorpAdapter {
	return &HashiCorpAdapter{Node: r, Address: a}
}

func (a *HashiCorpAdapter) NodeInfo(ctx context.Context) raftinspector.NodeInfo {
	id := a.Node.String()

	state := a.Node.State().String()

	term, err := strconv.ParseUint(a.Node.Stats()["term"], 10, 64)
	if err != nil {
		log.Printf("invalid term: %v", err)
		term = 0
	}

	commitIndex, err := strconv.ParseUint(a.Node.Stats()["commit_index"], 10, 64)
	if err != nil {
		log.Printf("invalid commit_index: %v", err)
		commitIndex = 0
	}

	appliedIndex, err := strconv.ParseUint(a.Node.Stats()["applied_index"], 10, 64)
	if err != nil {
		log.Printf("invalid commit_index: %v", err)
		appliedIndex = 0
	}

	lastSnapshotIndex, err := strconv.ParseUint(a.Node.Stats()["last_snapshot_index"], 10, 64)
	if err != nil {
		log.Printf("invalid last_snapshot_index: %v", err)
		lastSnapshotIndex = 0
	}

	return raftinspector.NodeInfo{
		ID:                id,
		Address:           a.Address,
		State:             state,
		Term:              term,
		CommitIndex:       commitIndex,
		AppliedIndex:      appliedIndex,
		LastSnapshotIndex: lastSnapshotIndex,
	}
}

func (a *HashiCorpAdapter) ClusterInfo(ctx context.Context) raftinspector.ClusterInfo {
	return raftinspector.ClusterInfo{}
}
