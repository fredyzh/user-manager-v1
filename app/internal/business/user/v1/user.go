package v1

import (
	"context"
	"strconv"
	"strings"

	pb "github.com/fredyzh/user_manager/api/user/v1/pb"
	pbs "github.com/fredyzh/user_manager/api/user/v1/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc UserUseCase) CreateUser(ctx context.Context, req *pbs.CreateUserReq) (*pbs.CreateUserReply, error) {
	hash, err := HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	objID, err := uc.DB.CreateUser(ctx, &pb.User{
		BasicAuth: &pb.BasicAuth{
			Username: req.Username,
			Password: hash,
		},
		Scope: req.Scope,
	})

	if err != nil {
		return nil, err
	}

	return &pbs.CreateUserReply{
		Id: (objID.(primitive.ObjectID)).String(),
	}, nil
}

func (uc UserUseCase) GetUser(ctx context.Context, req *pbs.GetUserReq) (*pbs.GetUserReply, error) {
	usr, err := uc.DB.GetUser(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &pbs.GetUserReply{
		User: usr,
	}, nil
}

func (uc UserUseCase) ListUser(ctx context.Context, req *pbs.ListUserReq) (*pbs.ListUserReply, error) {
	usrs, err := uc.DB.ListUser(ctx)
	if err != nil {
		return nil, err
	}

	return &pbs.ListUserReply{
		Results: usrs,
	}, nil
}

func (uc UserUseCase) GetUserByUsername(ctx context.Context, req *pbs.GetUserByUsernameReq) (*pbs.GetUserByUsernameReply, error) {
	usr, objId, err := uc.DB.GetUserByUsername(ctx, req.GetUsername(), req.GetScope())
	if err != nil {
		return nil, err
	}

	return &pbs.GetUserByUsernameReply{
		Id:   objId,
		User: usr,
	}, nil
}

func (uc UserUseCase) VerifyPassword(ctx context.Context, req *pbs.VerifyPasswordReq) (*pbs.VerifyPasswordReply, error) {
	ba, objId, err := uc.DB.GetUserByUsername(ctx, req.Username, req.Scope)

	if err != nil {
		if strings.Contains(err.Error(), "no doc") {
			return &pbs.VerifyPasswordReply{
				Ok: false,
			}, nil
		} else {
			return nil, err
		}
	}

	ok := CheckPasswordHash(req.Password, ba.BasicAuth.Password)

	return &pbs.VerifyPasswordReply{
		Id: objId,
		Ok: ok,
	}, nil
}

func (uc UserUseCase) SaveAddress(ctx context.Context, req *pbs.SaveAddressReq) (*pbs.SaveAddressReply, error) {
	r, err := uc.DB.SaveAddress(ctx, req.Id, req.Address)
	if err != nil {
		return nil, err
	}

	saveCnt := strconv.FormatInt(r.(int64), 10)
	return &pbs.SaveAddressReply{
		Id: saveCnt,
	}, nil
}

func (uc UserUseCase) SaveProfile(ctx context.Context, req *pbs.SaveProfileReq) (*pbs.SaveProfileReply, error) {
	r, err := uc.DB.SaveProfile(ctx, req.Id, req.Profile)
	if err != nil {
		return nil, err
	}

	saveCnt := strconv.FormatInt(r.(int64), 10)
	return &pbs.SaveProfileReply{
		Id: saveCnt,
	}, nil
}
