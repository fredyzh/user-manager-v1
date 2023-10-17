package v1

import (
	"context"
	"log/slog"

	gw1 "github.com/fredyzh/user_manager/api/token/v1/service"
	gw "github.com/fredyzh/user_manager/api/user/v1/service"
	conf "github.com/fredyzh/user_manager/app/configs/v1"
	svc "github.com/fredyzh/user_manager/app/internal/service/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func NewHTTPServer(c *conf.Server_HTTP, s *svc.UserService) *runtime.ServeMux {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, c.Timeout.AsDuration())
	defer cancel()

	mux := runtime.NewServeMux()
	err := gw.RegisterUserServiceHandlerServer(ctx, mux, s)
	if err != nil {
		slog.Error("err register token http server: %v\n", err)
		return nil
	}

	return mux
}

func NewHTTPTokenServer(c *conf.Server_HTTP, s *svc.TokenService) *runtime.ServeMux {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, c.Timeout.AsDuration())
	defer cancel()

	mux := runtime.NewServeMux()
	err := gw1.RegisterTokenServiceHandlerServer(ctx, mux, s)
	if err != nil {
		slog.Error("err register token http server: %v\n", err)
		return nil
	}

	return mux
}
