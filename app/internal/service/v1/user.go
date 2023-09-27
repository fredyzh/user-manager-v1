package v1

import (
	"context"
	"log/slog"

	pb "github.com/fredyzh/user_manager/api/user/v1/service"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserReply, error) {
	rv, err := s.UB.CreateUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserReply{
		Id: rv.Id,
	}, nil
}

func (s *UserService) Health(ctx context.Context, e *emptypb.Empty) (*pb.HealthReply, error) {
	slog.Debug("health called")
	return &pb.HealthReply{
		Ok: true,
	}, nil
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserReq) (*pb.GetUserReply, error) {

	r, err := s.UB.GetUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return r, nil
}
func (s *UserService) GetUserByUsername(ctx context.Context, req *pb.GetUserByUsernameReq) (*pb.GetUserByUsernameReply, error) {
	r, err := s.UB.GetUserByUsername(ctx, req)
	if err != nil {
		return nil, err
	}
	return r, nil
}
func (s *UserService) ResetUserPassword(ctx context.Context, req *pb.ResetUserPasswordReq) (*pb.ResetUserPasswordReply, error) {
	return &pb.ResetUserPasswordReply{}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserReq) (*pb.ListUserReply, error) {
	r, err := s.UB.ListUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return r, nil
}
func (s *UserService) VerifyPassword(ctx context.Context, req *pb.VerifyPasswordReq) (*pb.VerifyPasswordReply, error) {
	slog.Debug(req.String())
	return &pb.VerifyPasswordReply{}, nil
}
func (s *UserService) SaveAddress(ctx context.Context, req *pb.SaveAddressReq) (*pb.SaveAddressReply, error) {
	slog.Debug(req.String())
	return &pb.SaveAddressReply{}, nil
}
func (s *UserService) GetUserAddress(ctx context.Context, req *pb.GetAddressReq) (*pb.GetAddressReply, error) {
	slog.Debug(req.String())
	return &pb.GetAddressReply{}, nil
}

func (s *UserService) SaveProfile(ctx context.Context, req *pb.SaveProfileReq) (*pb.SaveProfileReply, error) {
	slog.Debug(req.String())
	return &pb.SaveProfileReply{}, nil
}
func (s *UserService) GetUserProfile(ctx context.Context, req *pb.GetProfileReq) (*pb.GetProfileReply, error) {
	slog.Debug(req.String())
	return &pb.GetProfileReply{}, nil
}
