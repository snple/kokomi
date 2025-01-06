package core

import (
	"context"

	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/cores"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthService struct {
	cs *CoreService

	cores.UnimplementedAuthServiceServer
}

func newAuthService(cs *CoreService) *AuthService {
	return &AuthService{
		cs: cs,
	}
}

func (s *AuthService) Login(ctx context.Context, in *cores.LoginRequest) (*cores.LoginResponse, error) {
	var err error
	var output cores.LoginResponse

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if in.GetName() == "" {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid User")
		}

		if len(in.GetPass()) < 8 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid Pass")
		}
	}

	item, err := s.cs.GetUser().ViewByName(ctx, in.GetName())
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(item.Pass), []byte(in.GetPass()))
	if err != nil {
		return &output, status.Errorf(codes.Unauthenticated, "Invalid Pass: %v", err)
	}

	outputUser := pb.User{}
	s.cs.GetUser().copyModelToOutput(&outputUser, &item)

	output.User = &outputUser

	return &output, nil
}

func (s *AuthService) ChangePass(ctx context.Context, in *cores.ChangePassRequest) (*pb.Message, error) {
	var err error
	var output pb.Message
	output.Message = "Failed"

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetOldPass()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid current Pass")
		}

		if len(in.GetNewPass()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid new Pass")
		}

		if len(in.GetNewPass()) < 8 {
			return &output, status.Error(codes.InvalidArgument, "Pass min 8 character")
		}
	}

	item, err := s.cs.GetUser().ViewByID(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(item.Pass), []byte(in.GetOldPass()))
	if err != nil {
		return &output, status.Errorf(codes.Unauthenticated, "Invalid Pass: %v", err)
	}

	passwd, err := bcrypt.GenerateFromPassword([]byte(in.GetNewPass()), bcrypt.DefaultCost)
	if err != nil {
		return &output, status.Errorf(codes.Internal, "Hash Pass: %v", err)
	}

	item.Pass = string(passwd)

	_, err = s.cs.GetDB().NewUpdate().Model(&item).Column("pass").WherePK().Exec(ctx)
	if err != nil {
		return &output, status.Errorf(codes.Internal, "Update: %v", err)
	}

	output.Message = "Success"

	return &output, nil
}

func (s *AuthService) ForceChangePass(ctx context.Context, in *cores.ForceChangePassRequest) (*pb.Message, error) {
	var err error
	var output pb.Message
	output.Message = "Failed"

	// basic validation
	{
		if in == nil {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid argument")
		}

		if len(in.GetNewPass()) == 0 {
			return &output, status.Error(codes.InvalidArgument, "Please supply valid new Pass")
		}

		if len(in.GetNewPass()) < 8 {
			return &output, status.Error(codes.InvalidArgument, "Pass min 8 character")
		}
	}

	item, err := s.cs.GetUser().ViewByID(ctx, in.GetId())
	if err != nil {
		return &output, err
	}

	passwd, err := bcrypt.GenerateFromPassword([]byte(in.GetNewPass()), bcrypt.DefaultCost)
	if err != nil {
		return &output, status.Errorf(codes.Internal, "Hash Pass: %v", err)
	}

	item.Pass = string(passwd)

	_, err = s.cs.GetDB().NewUpdate().Model(&item).Column("pass").WherePK().Exec(ctx)
	if err != nil {
		return &output, status.Errorf(codes.Internal, "Update: %v", err)
	}

	output.Message = "Success"

	return &output, nil
}
