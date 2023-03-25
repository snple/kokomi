package slot

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"snple.com/kokomi/pb"
	"snple.com/kokomi/pb/edges"
	"snple.com/kokomi/pb/slots"
)

type DataService struct {
	ss *SlotService

	slots.UnimplementedDataServiceServer
}

func newDataService(ss *SlotService) *DataService {
	return &DataService{
		ss: ss,
	}
}

func (s *DataService) Upload(ctx context.Context, in *slots.DataUploadRequest) (*slots.DataUploadResponse, error) {
	var err error
	var output slots.DataUploadResponse
	output.Message = "Failed"

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

	request := &edges.DataUploadRequest{
		Id:          in.GetId(),
		ContentType: in.GetContentType(),
		Content:     in.GetContent(),
		Cache:       in.GetCache(),
	}

	reply, err := s.ss.es.GetData().Upload(ctx, request)
	if err != nil {
		return nil, err
	}

	output.Id = reply.GetId()
	output.Message = reply.GetMessage()

	return &output, nil
}

func (s *DataService) Compile(ctx context.Context, in *slots.DataQueryRequest) (*pb.Message, error) {
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

	request := &edges.DataQueryRequest{
		Flux: in.GetFlux(),
		Vars: in.GetVars(),
	}

	return s.ss.es.GetData().Compile(ctx, request)
}

func (s *DataService) Query(in *slots.DataQueryRequest, stream slots.DataService_QueryServer) error {
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

	request := &edges.DataQueryRequest{
		Flux: in.GetFlux(),
		Vars: in.GetVars(),
	}

	return s.ss.es.GetData().Query(request, stream)
}

func (s *DataService) QueryTag(in *slots.DataQueryTagRequest, stream slots.DataService_QueryTagServer) error {
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

	request := &edges.DataQueryTagRequest{
		Id:   in.GetId(),
		Vars: in.GetVars(),
	}

	return s.ss.es.GetData().QueryTag(request, stream)
}
