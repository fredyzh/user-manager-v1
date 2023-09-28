package v1

import (
	"context"

	pbt "github.com/fredyzh/user_manager/api/token/v1/pb"
	pbu "github.com/fredyzh/user_manager/api/token/v1/service"
	pb "github.com/fredyzh/user_manager/api/user/v1/service"
	db "github.com/fredyzh/user_manager/app/internal/data/user/v1"
)

type UserUseCase struct {
	DB               db.DataRepo
	AuthKey          string
	JwtIssuer        string
	JwtTokenExpire   int64
	JwtRefreshExpire int64
}

type UserBusiness interface {
	GetUser(ctx context.Context, req *pb.GetUserReq) (*pb.GetUserReply, error)
	GetUserByUsername(ctx context.Context, req *pb.GetUserByUsernameReq) (*pb.GetUserByUsernameReply, error)
	// ResetUserPassword(ctx context.Context, req *pb.ResetUserPasswordReq) (*pb.ResetUserPasswordReply, error)
	CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserReply, error)
	ListUser(ctx context.Context, req *pb.ListUserReq) (*pb.ListUserReply, error)
	VerifyPassword(ctx context.Context, req *pb.VerifyPasswordReq) (*pb.VerifyPasswordReply, error)
	// GetUserAddress(ctx context.Context, req *pb.GetAddressReq) (*pb.GetAddressReply, error)
	SaveProfile(ctx context.Context, req *pb.SaveProfileReq) (*pb.SaveProfileReply, error)
	// GetUserProfile(ctx context.Context, req *pb.GetProfileReq) (*pb.GetProfileReply, error)
	SaveAddress(ctx context.Context, req *pb.SaveAddressReq) (*pb.SaveAddressReply, error)
	RegisterJWTSecret(ctx context.Context, req *pbu.RegisterJWTSecretRequest) (*pbu.RegisterJWTSecretReply, error)
	IssueJWTToken(ctx context.Context, req *pbu.IssueTokenRequest) (*pbu.IssueTokenReply, error)
	ValidateToken(ctx context.Context, req *pbu.ValidateTokenRequest) (*pbu.ValidateTokenReply, error)
	VerifyCode(ctx context.Context, token *pbu.VerifyCodeRequest) (*pbt.Code, error)
	RefreshToken(ctx context.Context, req *pbu.RefreshTokenRequest) (*pbu.RefreshTokenReply, error)
}
