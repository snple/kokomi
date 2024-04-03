package slot

import (
	"context"

	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/edges"
	"github.com/snple/kokomi/pb/slots"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		Realtime:    in.GetRealtime(),
		Save:        in.GetSave(),
	}

	reply, err := s.ss.Edge().GetData().Upload(ctx, request)
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

func (s *DataService) QueryById(in *slots.DataQueryByIdRequest, stream slots.DataService_QueryByIdServer) error {
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

	request := &edges.DataQueryByIdRequest{
		Id:   in.GetId(),
		Vars: in.GetVars(),
	}

	return s.ss.es.GetData().QueryById(request, stream)
}
