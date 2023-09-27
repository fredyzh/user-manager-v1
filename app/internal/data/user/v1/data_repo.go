package v1

import (
	"context"

	meta "github.com/fredyzh/user_manager/api/meta/v1/pb"
	pbt "github.com/fredyzh/user_manager/api/token/v1/pb"
	pb "github.com/fredyzh/user_manager/api/user/v1/pb"
)

type DataRepo interface {
	ConnectDB() interface{}
	CreateUser(ctx context.Context, usr *pb.User) (interface{}, error)
	GetUser(ctx context.Context, id string) (*pb.User, error)
	GetUserByUsername(ctx context.Context, name string, scope *meta.Scope) (*pb.User, string, error)
	// ResetUserPassword(id, opw, npw string) error
	ListUser(ctx context.Context) ([]*pb.User, error)

	SaveAddress(ctx context.Context, id string, addr *pb.Address) (interface{}, error)
	SaveProfile(ctx context.Context, id string, prof *pb.Profile) (interface{}, error)
	SaveJWTSecret(ctx context.Context, token *pbt.JWTSecrets) (interface{}, error)
	GetJwtSecret(ctx context.Context, userId string, key string) (string, error)
	// VerifyCode(ctx context.Context, token *pbt.VerifyCodeRequest) (*pbt.Code, error)
}
