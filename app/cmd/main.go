package main

import (
	"flag"
	"log/slog"
	"os"

	conf "github.com/fredyzh/user_manager/app/configs/v1"
	bus "github.com/fredyzh/user_manager/app/internal/business/user/v1"
	svc "github.com/fredyzh/user_manager/app/internal/service/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"

	"sigs.k8s.io/yaml"
)

var (
	// Name is the name of the compiled software.
	Name = "user.Manager"
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "./app/configs/config.yaml", "config path, eg: -conf config.yaml")
}

func newApp(hs *runtime.ServeMux, gs *grpc.Server, cg *conf.Server_GRPC, ch *conf.Server_HTTP, au *conf.Auth, ub *bus.UserBusiness) *App {
	return &App{
		grpcServer:    gs,
		grpcConfig:    cg,
		httpServerMux: hs,
		httpConfig:    ch,
		userBus:       ub,
		auth:          au,
	}
}

func newTokenApp(hs *runtime.ServeMux, gs *grpc.Server, cg *conf.Server_GRPC, ch *conf.Server_HTTP, au *conf.Auth, ub *bus.UserBusiness) *App {
	return &App{
		grpcServer:    gs,
		grpcConfig:    cg,
		httpServerMux: hs,
		httpConfig:    ch,
		userBus:       ub,
		auth:          au,
	}
}

func main() {
	flag.Parse()

	programLevel := new(slog.LevelVar) // Info by default
	programLevel.Set(slog.LevelDebug)
	h := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel})
	slog.SetDefault(slog.New(h))

	yamlFile, err := os.ReadFile(flagconf)
	if err != nil {
		slog.Error("Error reading YAML file: %s\n", err)
		return
	}

	jsn, err := yaml.YAMLToJSON(yamlFile)
	if err != nil {
		slog.Error("err yaml to json: %v\n", err)
		return
	}

	var bc conf.Bootstrap = conf.Bootstrap{}
	err = protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(jsn, &bc)
	if err != nil {
		slog.Error("err parse json: %v\n", err)
		return
	}

	app, err := initApp(bc.Server.Grpc, bc.Server.Http, bc.Data.Database, bc.Data.Redis, bc.Auth, &svc.UserService{}, &svc.TokenService{})
	if err != nil {
		slog.Error("init app failed: %v\n", err)
		return
	}

	app1, err := initTokenApp(bc.Server.Grpc, bc.Server.Http, bc.Data.Database, bc.Data.Redis, bc.Auth, &svc.UserService{}, &svc.TokenService{})
	if err != nil {
		slog.Error("init token app failed: %v\n", err)
		return
	}

	go app.StartHttp(app.httpConfig.Port)
	go app.StartGrpc(app.grpcConfig.Port)

	go app1.StartHttp(app1.httpConfig.Port + 1)
	app1.StartGrpc(app1.grpcConfig.Port + 1)
}
