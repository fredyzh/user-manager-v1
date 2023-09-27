package main

import (
	"fmt"
	"log/slog"
	"net"
	"net/http"

	conf "github.com/fredyzh/user_manager/app/configs/v1"
	bus "github.com/fredyzh/user_manager/app/internal/business/user/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type App struct {
	grpcServer         *grpc.Server
	grpcTokenServer    *grpc.Server
	grpcConfig         *conf.Server_GRPC
	httpServerMux      *runtime.ServeMux
	httpTokenServerMux *runtime.ServeMux
	httpConfig         *conf.Server_HTTP
	userBus            *bus.UserBusiness
	auth               *conf.Auth
}

func (a *App) StartGrpc(port int64) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", a.grpcConfig.Addr, port))
	if err != nil {
		slog.Error("grpc listen failed", err)
		return
	}

	slog.Info(fmt.Sprintf("grpc server started at port: %s:%d", a.grpcConfig.Addr, port))

	if err = a.grpcServer.Serve(lis); err != nil {
		slog.Error("server failed", err)
		return
	}
}

func (a *App) StartHttp(port int64) {
	slog.Info(fmt.Sprintf("http server started at port: %s:%d", a.httpConfig.Addr, port))
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", a.httpConfig.Addr, port), a.httpServerMux)
	if err != nil {
		slog.Error("http failed", err)
		return
	}
}
