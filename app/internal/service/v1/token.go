package v1

import (
	"context"
	"log/slog"

	pb "github.com/fredyzh/user_manager/api/token/v1/service"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (t *TokenService) RegisterJWTSecret(ctx context.Context, req *pb.RegisterJWTSecretRequest) (*pb.RegisterJWTSecretReply, error) {
	rv, err := t.UB.RegisterJWTSecret(ctx, req)
	if err != nil {
		return nil, err
	}
	return rv, nil
}

func (t *TokenService) IssueJWTToken(ctx context.Context, req *pb.IssueTokenRequest) (*pb.IssueTokenReply, error) {
	rv, err := t.UB.IssueJWTToken(ctx, req)
	if err != nil {
		return nil, err
	}
	return rv, nil
}

func (t *TokenService) TokenHealth(ctx context.Context, e *emptypb.Empty) (*pb.TokenHealthReply, error) {
	slog.Debug("token health called")
	return &pb.TokenHealthReply{
		Ok: true,
	}, nil
}

func (t *TokenService) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenReply, error) {
	rv, err := t.UB.ValidateToken(ctx, req)
	if err != nil {
		return nil, err
	}
	return rv, nil
}

func (t *TokenService) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenReply, error) {
	rv, err := t.UB.RefreshToken(ctx, req)
	if err != nil {
		return nil, err
	}
	return rv, nil
}
