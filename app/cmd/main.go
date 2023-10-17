package main

import (
	"context"
	"flag"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	UDP = "udp"

	LOOK_UP_ADDR = "8.8.8.8:80"
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

	level := bc.Trace.Level
	programLevel := new(slog.LevelVar) // Info by default

	switch level {
	case conf.Loglevel_Debug:
		programLevel.Set(slog.LevelDebug)
	case conf.Loglevel_Error:
		programLevel.Set(slog.LevelError)
	case conf.Loglevel_Warn:
		programLevel.Set(slog.LevelWarn)
	}

	if bc.Server.Grpc.OnCloud {
		if bc.Server.Grpc.LookupHost {
			//cloud localhost
			local, remote, err := lookupHost()
			if err != nil {
				slog.Error(err.Error())
				return
			}
			if bc.Server.Grpc.UseRemote {
				bc.Server.Grpc.Addr = remote.IP.String()
			} else {
				bc.Server.Grpc.Addr = local.IP.String()
			}
		} else {
			bc.Server.Grpc.Addr = os.Getenv(bc.Server.Grpc.CloudEnv)
		}

		if len(bc.Server.Grpc.Addr) == 0 {
			slog.Error("grpc host not found.")
			return
		}
	}

	if bc.Server.Http.OnCloud {
		if bc.Server.Http.LookupHost {
			//cloud localhost
			local, remote, err := lookupHost()
			if err != nil {
				slog.Error(err.Error())
				return
			}
			if bc.Server.Http.UseRemote {
				bc.Server.Http.Addr = remote.IP.String()
			} else {
				bc.Server.Http.Addr = local.IP.String()
			}
		} else {
			bc.Server.Http.Addr = os.Getenv(bc.Server.Grpc.CloudEnv)
		}

		if len(bc.Server.Http.Addr) == 0 {
			slog.Error("http host not found.")
			return
		}
	}

	h := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel})
	slog.SetDefault(slog.New(h))

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
	go app1.StartGrpc(app1.grpcConfig.Port + 1)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	_, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	app.grpcServer.GracefulStop()
	app1.grpcServer.GracefulStop()

	slog.Info("Graceful shutdown complete.")
}

func lookupHost() (*net.UDPAddr, *net.UDPAddr, error) {
	conn, err := net.Dial(UDP, LOOK_UP_ADDR)
	if err != nil {
		slog.Error(err.Error())
		return nil, nil, err
	}
	defer conn.Close()

	localAddress := conn.LocalAddr().(*net.UDPAddr)
	remoteAddress := conn.RemoteAddr().(*net.UDPAddr)

	return localAddress, remoteAddress, nil
}
