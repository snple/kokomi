package slot

import (
	"context"

	"github.com/snple/beacon/pb"
	"github.com/snple/beacon/pb/slots"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NodeService struct {
	ss *SlotService

	slots.UnimplementedNodeServiceServer
}

func newNodeService(ss *SlotService) *NodeService {
	return &NodeService{
		ss: ss,
	}
}

func (s *NodeService) Update(ctx context.Context, in *pb.Node) (*pb.Node, error) {
	var output pb.Node
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	reply, err := s.ss.Edge().GetNode().View(ctx, &pb.MyEmpty{})
	if err != nil {
		return &output, err
	}

	in.Status = reply.GetStatus()

	return s.ss.Edge().GetNode().Update(ctx, in)
}

func (s *NodeService) View(ctx context.Context, in *pb.MyEmpty) (*pb.Node, error) {
	var output pb.Node
	var err error

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	_, err = validateToken(ctx)
	if err != nil {
		return &output, err
	}

	return s.ss.Edge().GetNode().View(ctx, in)
}
