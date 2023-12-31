// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/fredyzh/user_manager/app/configs/v1"
	v1_3 "github.com/fredyzh/user_manager/app/internal/server/v1"
	v1_2 "github.com/fredyzh/user_manager/app/internal/service/v1"
)

// Injectors from wire.go:

func initApp(server_GRPC *v1.Server_GRPC, server_HTTP *v1.Server_HTTP, data_Database *v1.Data_Database, data_Redis *v1.Data_Redis, auth *v1.Auth, userService *v1_2.UserService, tokenService *v1_2.TokenService) (*App, error) {
	serveMux := v1_3.NewHTTPServer(server_HTTP, userService)
	server := v1_3.NewGRPCServer(server_GRPC, userService)
	userBusiness := v1_3.NewDBInterface(data_Database, auth, userService)
	app := newApp(serveMux, server, server_GRPC, server_HTTP, auth, userBusiness)
	return app, nil
}

func initTokenApp(server_GRPC *v1.Server_GRPC, server_HTTP *v1.Server_HTTP, data_Database *v1.Data_Database, data_Redis *v1.Data_Redis, auth *v1.Auth, userService *v1_2.UserService, tokenService *v1_2.TokenService) (*App, error) {
	serveMux := v1_3.NewHTTPTokenServer(server_HTTP, tokenService)
	server := v1_3.NewGRPCTokenServer(server_GRPC, tokenService)
	userBusiness := v1_3.NewTokenDBInterface(data_Database, auth, tokenService)
	app := newTokenApp(serveMux, server, server_GRPC, server_HTTP, auth, userBusiness)
	return app, nil
}
