package v1

import (
	token "github.com/fredyzh/user_manager/api/token/v1/service"
	svr "github.com/fredyzh/user_manager/api/user/v1/service"
	bus "github.com/fredyzh/user_manager/app/internal/business/user/v1"
)

type UserService struct {
	svr.UnimplementedUserServiceServer

	UB bus.UserBusiness
}

type TokenService struct {
	token.UnimplementedTokenServiceServer

	UB bus.UserBusiness
}
