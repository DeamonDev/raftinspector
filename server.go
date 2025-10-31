package raftinspector

import (
	"context"
	"net"

	"github.com/deamondev/raftinspector/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	proto.UnimplementedRaftInspectorServer
	adapter Adapter
}

func NewServer(adapter Adapter) *Server {
	return &Server{adapter: adapter}
}

func (s *Server) GetNodeInfo(ctx context.Context, _ *emptypb.Empty) (*proto.NodeInfo, error) {
	nodeInfo := s.adapter.GetNodeInfo(ctx)

	return &proto.NodeInfo{
		Id:                nodeInfo.ID,
		Address:           nodeInfo.Address,
		State:             nodeInfo.State,
		Term:              nodeInfo.Term,
		CommitIndex:       nodeInfo.CommitIndex,
		AppliedIndex:      nodeInfo.AppliedIndex,
		LastSnapshotIndex: nodeInfo.LastSnapshotIndex,
	}, nil
}

func (s *Server) GetClusterInfo(ctx context.Context, _ *emptypb.Empty) (*proto.ClusterInfo, error) {
	return nil, nil
}

func (s *Server) Run(network string, address string) error {
	l, _ := net.Listen(network, address) // example being [network=tcp, address=:50051]
	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	proto.RegisterRaftInspectorServer(grpcServer, s)

	return grpcServer.Serve(l)
}
