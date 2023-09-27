package v1

import (
	"context"

	pb "github.com/fredyzh/user_manager/api/token/v1/service"
)

func (t *TokenService) RegisterJWTSecret(ctx context.Context, req *pb.RegisterJWTSecretRequest) (*pb.RegisterJWTSecretReply, error) {
	rv, err := t.UB.RegisterJWTSecret(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.RegisterJWTSecretReply{
		Id: rv.Id,
	}, nil
}

func (t *TokenService) IssueJWTToken(ctx context.Context, req *pb.IssueTokenRequest) (*pb.IssueTokenReply, error) {
	rv, err := t.UB.IssueJWTToken(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.IssueTokenReply{
		Id:          rv.Id,
		AccessToken: rv.AccessToken,
	}, nil
}
