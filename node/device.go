package node

import (
	"context"

	"github.com/snple/kokomi/consts"
	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/cores"
	"github.com/snple/kokomi/pb/nodes"
	"github.com/snple/kokomi/util/metadata"
	"github.com/snple/kokomi/util/token"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DeviceService struct {
	ns *NodeService

	nodes.UnimplementedDeviceServiceServer
}

func newDeviceService(ns *NodeService) *DeviceService {
	return &DeviceService{
		ns: ns,
	}
}

func (s *DeviceService) Login(ctx context.Context, in *nodes.DeviceLoginRequest) (*nodes.DeviceLoginReply, error) {
	var output nodes.DeviceLoginReply

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetId() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Device.ID")
		}

		if len(in.GetSecret()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Device.Secret")
		}
	}

	request := &pb.Id{Id: in.GetId()}

	reply, err := s.ns.Core().GetDevice().View(ctx, request)
	if err != nil {
		return &output, err
	}

	if reply.GetStatus() != consts.ON {
		s.ns.Logger().Sugar().Errorf("device connect error: device is not enable, id: %v, ip: %v",
			in.GetId(), metadata.GetPeerAddr(ctx))
		return &output, status.Error(codes.FailedPrecondition, "The device is not enable")
	}

	if reply.GetSecret() != string(in.GetSecret()) {
		s.ns.Logger().Sugar().Errorf("device connect error: device secret is not valid, id: %v, ip: %v",
			in.GetId(), metadata.GetPeerAddr(ctx))
		return &output, status.Error(codes.Unauthenticated, "Please supply valid secret")
	}

	token, err := token.ClaimDeviceToken(reply.Id)
	if err != nil {
		return &output, status.Errorf(codes.Internal, "claim token: %v", err)
	}

	s.ns.Logger().Sugar().Infof("device connect success, id: %v, ip: %v", in.GetId(), metadata.GetPeerAddr(ctx))

	reply.Secret = ""

	output.Device = reply
	output.Token = token

	return &output, nil
}

func validateToken(ctx context.Context) (deviceID string, err error) {
	tks, err := metadata.GetToken(ctx)
	if err != nil {
		return "", err
	}

	ok, deviceID := token.ValidateDeviceToken(tks)
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "Token validation failed")
	}

	return deviceID, nil
}

func (s *DeviceService) Update(ctx context.Context, in *pb.Device) (*pb.Device, error) {
	var output pb.Device
	// var err error

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

	if in.GetId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: in.GetId() != deviceID")
	}

	request := &pb.Id{Id: deviceID}

	reply, err := s.ns.Core().GetDevice().View(ctx, request)
	if err != nil {
		return &output, err
	}

	in.Secret = reply.GetSecret()
	in.Status = reply.GetStatus()

	return s.ns.Core().GetDevice().Update(ctx, in)
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

	deviceID, err := validateToken(ctx)
	if err != nil {
		return &output, err
	}

	request := &pb.Id{Id: deviceID}

	reply, err := s.ns.Core().GetDevice().View(ctx, request)
	if err != nil {
		return &output, err
	}

	reply.Secret = ""

	return reply, err
}

func (s *DeviceService) Link(ctx context.Context, in *nodes.DeviceLinkRequest) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

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

	request := &cores.DeviceLinkRequest{Id: deviceID, Status: in.GetStatus()}

	return s.ns.Core().GetDevice().Link(ctx, request)
}

func (s *DeviceService) ViewWithDeleted(ctx context.Context, in *pb.MyEmpty) (*pb.Device, error) {
	var output pb.Device
	var err error

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

	request := &pb.Id{Id: deviceID}

	reply, err := s.ns.Core().GetDevice().ViewWithDeleted(ctx, request)
	if err != nil {
		return &output, err
	}

	reply.Secret = ""

	return reply, err
}

func (s *DeviceService) Sync(ctx context.Context, in *pb.Device) (*pb.MyBool, error) {
	var err error
	var output pb.MyBool

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

	if in.GetId() != deviceID {
		return &output, status.Error(codes.NotFound, "Query: in.GetId() != deviceID")
	}

	request := &pb.Id{Id: deviceID}

	reply, err := s.ns.Core().GetDevice().ViewWithDeleted(ctx, request)
	if err != nil {
		return &output, err
	}

	in.Secret = reply.GetSecret()
	in.Status = reply.GetStatus()
	in.Deleted = reply.GetDeleted()

	return s.ns.Core().GetDevice().Sync(ctx, in)
}

func (s *DeviceService) KeepAlive(in *pb.MyEmpty, stream nodes.DeviceService_KeepAliveServer) error {
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

	<-stream.Context().Done()

	return stream.Context().Err()
}
