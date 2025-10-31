package hashicorp

import (
	"context"

	"github.com/deamondev/raftinspector"
	"github.com/hashicorp/raft"
)

/*
	HashiCorpAdapter Use this adapter if you are able to instrument your raft nodes.

If you're not able to modify code, then you should use a metrics server adapter
In that case your adapter will be readonly.
This adapter allows modifying the state of the cluster
*/
type HashiCorpAdapter struct {
	Node    *raft.Raft
	NodeID  string
	Address string
}

func NewHashiCorpAdapter(r *raft.Raft, nodeID string, a string) *HashiCorpAdapter {
	return &HashiCorpAdapter{Node: r, NodeID: nodeID, Address: a}
}

func (a *HashiCorpAdapter) GetNodeInfo(ctx context.Context) raftinspector.NodeInfo {
	nodeID := a.NodeID

	state := a.Node.State().String()

	term := a.Node.Stats()["term"]
	if term == "" {
		term = "undefined"
	}

	commitIndex := a.Node.Stats()["commit_index"]
	if commitIndex == "" {
		term = "undefined"
	}

	appliedIndex := a.Node.Stats()["applied_index"]
	if appliedIndex == "" {
		term = "undefined"
	}

	lastSnapshotIndex := a.Node.Stats()["last_snapshot_index"]
	if lastSnapshotIndex == "" {
		term = "undefined"
	}

	return raftinspector.NodeInfo{
		ID:                nodeID,
		Address:           a.Address,
		State:             state,
		Term:              term,
		CommitIndex:       commitIndex,
		AppliedIndex:      appliedIndex,
		LastSnapshotIndex: lastSnapshotIndex,
	}
}

func (a *HashiCorpAdapter) GetClusterInfo(ctx context.Context) raftinspector.ClusterInfo {
	return raftinspector.ClusterInfo{}
}
