package node

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"snple.com/kokomi/pb"
	"snple.com/kokomi/pb/cores"
	"snple.com/kokomi/pb/nodes"
)

type DataService struct {
	ns *NodeService

	nodes.UnimplementedDataServiceServer
}

func newDataService(ns *NodeService) *DataService {
	return &DataService{
		ns: ns,
	}
}

func (s *DataService) Upload(ctx context.Context, in *nodes.DataUploadRequest) (*nodes.DataUploadResponse, error) {
	var err error
	var output nodes.DataUploadResponse
	output.Message = "Failed"

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	request := &cores.DataUploadRequest{
		Id:          in.GetId(),
		ContentType: in.GetContentType(),
		Content:     in.GetContent(),
		DeviceId:    deviceID,
		Cache:       in.GetCache(),
		Save:        in.GetSave(),
	}

	reply, err := s.ns.cs.GetData().Upload(ctx, request)
	if err != nil {
		return nil, err
	}

	output.Id = reply.GetId()
	output.Message = reply.GetMessage()

	return &output, nil
}

func (s *DataService) Compile(ctx context.Context, in *nodes.DataQueryRequest) (*pb.Message, error) {
	var err error
	var output pb.Message

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

	request := &cores.DataQueryRequest{
		Flux: in.GetFlux(),
		Vars: in.GetVars(),
	}

	return s.ns.cs.GetData().Compile(ctx, request)
}

func (s *DataService) Query(in *nodes.DataQueryRequest, stream nodes.DataService_QueryServer) error {
	var err error

	// basic validation
	{
		if in == nil {
			return status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	_, err = validateToken(stream.Context())
	if err != nil {
		return err
	}

	request := &cores.DataQueryRequest{
		Flux: in.GetFlux(),
		Vars: in.GetVars(),
	}

	return s.ns.cs.GetData().Query(request, stream)
}

func (s *DataService) QueryTag(in *nodes.DataQueryTagRequest, stream nodes.DataService_QueryTagServer) error {
	var err error

	// basic validation
	{
		if in == nil {
			return status.Error(codes.InvalidArgument, "Please supply valid argument")
		}
	}

	_, err = validateToken(stream.Context())
	if err != nil {
		return err
	}

	request := &cores.DataQueryTagRequest{
		Id:   in.GetId(),
		Vars: in.GetVars(),
	}

	return s.ns.cs.GetData().QueryTag(request, stream)
}
