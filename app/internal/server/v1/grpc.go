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
		opts = append(opts, grpc.MaxRecvMsgSize(int(c.MaxRecvMsgSize)))
		opts = append(opts, grpc.MaxSendMsgSize(int(c.MaxSendMsgSize)))
		opts = append(opts, grpc.MaxConcurrentStreams(uint32(c.MaxConcurrentStreams)))
	}

	svr := grpc.NewServer(opts...)
	pb.RegisterUserServiceServer(svr, s)

	return svr
}

func NewGRPCTokenServer(c *conf.Server_GRPC, s *svc.TokenService) *grpc.Server {
	opts := []grpc.ServerOption{}
	if c.Timeout != nil {
		opts = append(opts, grpc.ConnectionTimeout(c.Timeout.AsDuration()))
	}

	svr := grpc.NewServer(opts...)
	pb1.RegisterTokenServiceServer(svr, s)

	return svr
}
