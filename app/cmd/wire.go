//go:build wireinject
// +build wireinject

package main

import (
	conf "github.com/fredyzh/user_manager/app/configs/v1"
	svr "github.com/fredyzh/user_manager/app/internal/server/v1"
	svc "github.com/fredyzh/user_manager/app/internal/service/v1"

	"github.com/google/wire"
)

func initApp(*conf.Server_GRPC, *conf.Server_HTTP, *conf.Data_Database, *conf.Data_Redis, *conf.Auth, *svc.UserService, *svc.TokenService) (*App, error) {
	panic(wire.Build(svr.ProviderSet, newApp))
}

func initTokenApp(*conf.Server_GRPC, *conf.Server_HTTP, *conf.Data_Database, *conf.Data_Redis, *conf.Auth, *svc.UserService, *svc.TokenService) (*App, error) {
	panic(wire.Build(svr.TokenProviderSet, newTokenApp))
}
