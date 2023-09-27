package v1

import (
	"context"
	"errors"
	"log"

	pbt "github.com/fredyzh/user_manager/api/token/v1/pb"
	pbu "github.com/fredyzh/user_manager/api/token/v1/service"
	pbs "github.com/fredyzh/user_manager/api/user/v1/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc UserUseCase) RegisterJWTSecret(ctx context.Context, req *pbu.RegisterJWTSecretRequest) (*pbu.RegisterJWTSecretReply, error) {
	var userId string
	if len(req.UserId) == 0 {
		rp, err := uc.VerifyPassword(ctx, &pbs.VerifyPasswordReq{
			Username: req.Username,
			Password: req.Password,
			Scope:    req.Scope,
		})
		if err != nil {
			return nil, err
		}

		if !rp.Ok {
			return nil, errors.New("invalid user")
		}

		userId = rp.Id
	}

	jwt := req.Jwt
	jwt.UserId = userId
	cipherData, err := EncryptToString(jwt.Value, []byte(uc.AuthKey))
	jwt.Value = cipherData

	objID, err := uc.DB.SaveJWTSecret(ctx, jwt)
	if err != nil {
		return nil, err
	}

	return &pbu.RegisterJWTSecretReply{
		Id: (objID.(primitive.ObjectID)).String(),
	}, nil
}

func (uc UserUseCase) IssueJWTToken(ctx context.Context, req *pbu.IssueTokenRequest) (*pbu.IssueTokenReply, error) {
	rp, err := uc.VerifyPassword(ctx, &pbs.VerifyPasswordReq{
		Username: req.Username,
		Password: req.Password,
		Scope:    req.Scope,
	})
	if err != nil {
		return nil, err
	}

	if !rp.Ok {
		return nil, errors.New("invalid user")
	}

	v, err := uc.DB.GetJwtSecret(ctx, rp.Id, req.Jwtkey)
	if err != nil {
		return nil, err
	}

	secret, err := DecryptFromString(v, []byte(uc.AuthKey))
	if err != nil {
		return nil, err
	}

	token, refresh, err := GenerateTokenPair(rp.Id, secret, req.Username, req.Scope.Domain, req.Scope.AppId, uc)
	if err != nil {
		return nil, err
	}

	log.Println(refresh)

	return &pbu.IssueTokenReply{
		Id:           rp.Id,
		AccessToken:  token,
		RefreshToken: refresh,
	}, nil
}

func (uc UserUseCase) ValidateToken(ctx context.Context, req *pbu.ValidateTokenRequest) (*pbt.Token, error) {
	return nil, nil
}

func (uc UserUseCase) VerifyCode(ctx context.Context, token *pbu.VerifyCodeRequest) (*pbt.Code, error) {
	return nil, nil
}
