package slot

import (
	"context"

	"github.com/snple/beacon/pb"
	"github.com/snple/beacon/pb/slots"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DeviceService struct {
	ss *SlotService

	slots.UnimplementedDeviceServiceServer
}

func newDeviceService(ss *SlotService) *DeviceService {
	return &DeviceService{
		ss: ss,
	}
}

func (s *DeviceService) Update(ctx context.Context, in *pb.Device) (*pb.Device, error) {
	var output pb.Device
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

	reply, err := s.ss.Edge().GetDevice().View(ctx, &pb.MyEmpty{})
	if err != nil {
		return &output, err
	}

	in.Status = reply.GetStatus()

	return s.ss.Edge().GetDevice().Update(ctx, in)
}

func (s *DeviceService) View(ctx context.Context, in *pb.MyEmpty) (*pb.Device, error) {
	var output pb.Device
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

	return s.ss.Edge().GetDevice().View(ctx, in)
}
