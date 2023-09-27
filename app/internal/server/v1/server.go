package v1

import (
	"strconv"

	conf "github.com/fredyzh/user_manager/app/configs/v1"
	bus "github.com/fredyzh/user_manager/app/internal/business/user/v1"
	mongo "github.com/fredyzh/user_manager/app/internal/data/user/v1/mongo"
	svc "github.com/fredyzh/user_manager/app/internal/service/v1"
	"github.com/google/wire"
	mgd "go.mongodb.org/mongo-driver/mongo"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewHTTPServer, NewGRPCServer, NewDBInterface)

var TokenProviderSet = wire.NewSet(NewHTTPTokenServer, NewGRPCTokenServer, NewTokenDBInterface)

func NewDBInterface(cb *conf.Data_Database, ca *conf.Auth, s *svc.UserService) *bus.UserBusiness {
	mongodb := &mongo.MongoDB{
		Host:      cb.Host,
		Port:      strconv.Itoa(int(cb.Port)),
		DefualtDb: cb.Defaultdb,
		UserDB:    cb.CollUser,
		TokenDB:   cb.CollToken,
		SecretsDB: cb.CollSecrets,
		Admin:     cb.User,
		Password:  cb.Password,
		Timeout:   cb.Timeout.AsDuration(),
	}

	mongodb.DBClient = mongodb.ConnectDB().(*mgd.Client)

	s.UB = bus.UserUseCase{DB: mongodb,
		AuthKey:          ca.Key,
		JwtIssuer:        ca.JwtIssuer,
		JwtTokenExpire:   ca.JwtTokenExpire,
		JwtRefreshExpire: ca.JwtRefreshExpire,
	}

	return &s.UB
}

func NewTokenDBInterface(cb *conf.Data_Database, ca *conf.Auth, s *svc.TokenService) *bus.UserBusiness {
	mongodb := &mongo.MongoDB{
		Host:      cb.Host,
		Port:      strconv.Itoa(int(cb.Port)),
		DefualtDb: cb.Defaultdb,
		UserDB:    cb.CollUser,
		TokenDB:   cb.CollToken,
		SecretsDB: cb.CollSecrets,
		Admin:     cb.User,
		Password:  cb.Password,
		Timeout:   cb.Timeout.AsDuration(),
	}

	mongodb.DBClient = mongodb.ConnectDB().(*mgd.Client)

	s.UB = bus.UserUseCase{DB: mongodb,
		AuthKey:          ca.Key,
		JwtIssuer:        ca.JwtIssuer,
		JwtTokenExpire:   ca.JwtTokenExpire,
		JwtRefreshExpire: ca.JwtRefreshExpire,
	}

	return &s.UB
}
