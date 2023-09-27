package v1

import (
	pb1 "github.com/fredyzh/user_manager/api/token/v1/service"
	pb "github.com/fredyzh/user_manager/api/user/v1/service"
	conf "github.com/fredyzh/user_manager/app/configs/v1"
	svc "github.com/fredyzh/user_manager/app/internal/service/v1"
	"google.golang.org/grpc"
)

func NewGRPCServer(c *conf.Server_GRPC, s *svc.UserService) *grpc.Server {
	opts := []grpc.ServerOption{}
	if c.Timeout != nil {
		opts = append(opts, grpc.ConnectionTimeout(c.Timeout.AsDuration()))
	}

	svr := grpc.NewServer(opts...)
	pb.RegisterUserServer(svr, s)

	return svr
}

func NewGRPCTokenServer(c *conf.Server_GRPC, s *svc.TokenService) *grpc.Server {
	opts := []grpc.ServerOption{}
	if c.Timeout != nil {
		opts = append(opts, grpc.ConnectionTimeout(c.Timeout.AsDuration()))
	}

	svr := grpc.NewServer(opts...)
	pb1.RegisterTokenServer(svr, s)

	return svr
}
