package v1

import (
	"context"
	"errors"
	"time"

	pbt "github.com/fredyzh/user_manager/api/token/v1/pb"
	pbu "github.com/fredyzh/user_manager/api/token/v1/service"
	pbs "github.com/fredyzh/user_manager/api/user/v1/service"
	"github.com/golang-jwt/jwt"
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

	return &pbu.IssueTokenReply{
		Id:           rp.Id,
		AccessToken:  token,
		RefreshToken: refresh,
	}, nil
}

func (uc UserUseCase) ValidateToken(ctx context.Context, req *pbu.ValidateTokenRequest) (*pbu.ValidateTokenReply, error) {

	tkn := ParseToken(req.AccessToken)

	claims := tkn.Claims.(jwt.MapClaims)

	usrId := claims["sub"].(string)

	isur := claims["iss"].(string)

	if isur != uc.JwtIssuer {
		return &pbu.ValidateTokenReply{
			IsValid:       false,
			InvalidReason: "invalid token",
		}, nil
	}

	secret, err := getJwtSecret(ctx, uc, usrId, req.Jwtkey)
	if err != nil {
		return &pbu.ValidateTokenReply{
			IsValid:       false,
			InvalidReason: "invalid token 1",
		}, nil
	}

	tkn, ep, err := ValidateToken(req.AccessToken, secret)
	if err != nil {
		return nil, err
	}

	if ep {
		return &pbu.ValidateTokenReply{
			IsValid:       false,
			InvalidReason: "expired",
		}, nil
	}

	return &pbu.ValidateTokenReply{
		IsValid: true,
	}, nil
}

func (uc UserUseCase) RefreshToken(ctx context.Context, req *pbu.RefreshTokenRequest) (*pbu.RefreshTokenReply, error) {
	old_tkn := ParseToken(req.AccessToken)
	old_claims := old_tkn.Claims.(jwt.MapClaims)
	ref_tkn := ParseToken(req.RefreshToken)
	ref_claims := ref_tkn.Claims.(jwt.MapClaims)

	old_isur := old_claims["iss"].(string)
	ref_isur := ref_claims["iss"].(string)

	if old_isur != ref_isur || old_isur != uc.JwtIssuer || ref_isur != uc.JwtIssuer {
		return &pbu.RefreshTokenReply{
			Ok:            false,
			FailureReason: "invalid tokens 2",
		}, nil
	}

	old_usr := old_claims["sub"].(string)
	ref_usr := ref_claims["sub"].(string)

	if old_usr != ref_usr {
		return &pbu.RefreshTokenReply{
			Ok:            false,
			FailureReason: "invalid tokens 3",
		}, nil
	}

	secret, err := getJwtSecret(ctx, uc, ref_usr, req.Jwtkey)
	if err != nil {
		return &pbu.RefreshTokenReply{
			Ok:            false,
			FailureReason: "invalid tokens 4",
		}, nil
	}

	_, ep, _ := ValidateToken(req.RefreshToken, secret)
	if ep {
		return &pbu.RefreshTokenReply{
			Ok:            false,
			FailureReason: "refresh tokens expired",
		}, nil
	}

	//set expriry for JWT
	old_claims["exp"] = time.Now().UTC().Add(time.Minute * time.Duration(uc.JwtTokenExpire)).Unix()
	old_claims["iat"] = time.Now().UTC().Unix()

	//create singed token
	signedAccessToken, err := old_tkn.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	return &pbu.RefreshTokenReply{
		AccessToken:  signedAccessToken,
		RefreshToken: req.RefreshToken,
		Ok:           true,
	}, nil
}

func (uc UserUseCase) VerifyCode(ctx context.Context, token *pbu.VerifyCodeRequest) (*pbt.Code, error) {
	return nil, nil
}

func getJwtSecret(ctx context.Context, uc UserUseCase, usrId, jwtkey string) (string, error) {
	v, err := uc.DB.GetJwtSecret(ctx, usrId, jwtkey)
	if err != nil {
		return "", err
	}

	secret, err := DecryptFromString(v, []byte(uc.AuthKey))
	if err != nil {
		return "", err
	}
	return secret, nil
}
