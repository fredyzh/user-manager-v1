package grpctest

import (
	"context"
	"log"
	"testing"

	pbt "github.com/fredyzh/user_manager/api/token/v1/service"
	pb "github.com/fredyzh/user_manager/api/user/v1/service"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

var addr string = "localhost:17777"
var addrtoken string = "localhost:17778"

func TestGRPCServerHealth(t *testing.T) {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.Dial(addr, opts...)

	if err != nil {
		t.Error(err)
		t.Failed()
	}

	defer conn.Close()

	clt := pb.NewUserClient(conn)

	ok, err := clt.Health(context.Background(), &emptypb.Empty{})

	if err != nil {
		t.Error(err)
		t.Failed()
	}

	if !ok.Ok {
		t.Failed()
	}
}

func TestGRPCTokenServerHealth(t *testing.T) {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.Dial(addrtoken, opts...)

	if err != nil {
		t.Error(err)
		t.Failed()
	}

	defer conn.Close()

	clt := pbt.NewTokenClient(conn)

	ok, err := clt.TokenHealth(context.Background(), &emptypb.Empty{})

	if err != nil {
		t.Error(err)
		t.Failed()
	}

	if !ok.Ok {
		t.Failed()
	}
}

func TestGRPCServerAPIs(t *testing.T) {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.Dial(addr, opts...)
	defer conn.Close()

	if err != nil {
		t.Error(err)
		t.Failed()
	}

	clt := pb.NewUserClient(conn)

	_, err = clt.GetUser(context.Background(), &TestGetUserReq)

	if err != nil {
		t.Error(err)
		t.Failed()
	}

	_, err = clt.CreateUser(context.Background(), &TestCreateUserReq)

	if err != nil {
		t.Error(err)
		t.Failed()
	}

	_, err = clt.GetUserByUsername(context.Background(), &TestGetUserByUsernameReq)
	if err != nil {
		t.Error(err)
		t.Failed()
	}

	_, err = clt.SaveProfile(context.Background(), &TestSaveProfileReq)
	if err != nil {
		t.Error(err)
		t.Failed()
	}

	_, err = clt.GetUserProfile(context.Background(), &TestGetProfileReq)
	if err != nil {
		t.Error(err)
		t.Failed()
	}

	_, err = clt.SaveAddress(context.Background(), &TestCreateAddressReq)
	if err != nil {
		t.Error(err)
		t.Failed()
	}

	_, err = clt.GetUserAddress(context.Background(), &TestGetAddressReq)
	if err != nil {
		t.Error(err)
		t.Failed()
	}

	_, err = clt.ResetUserPassword(context.Background(), &TestResetUserPasswordReq)
	if err != nil {
		t.Error(err)
		t.Failed()
	}

	_, err = clt.VerifyPassword(context.Background(), &TestVerifyPasswordReq)
	if err != nil {
		t.Error(err)
		t.Failed()
	}
}

func TestGRPCServerLists(t *testing.T) {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.Dial(addr, opts...)
	defer conn.Close()

	if err != nil {
		t.Error(err)
		t.Failed()
	}

	clt := pb.NewUserClient(conn)

	users, err := clt.ListUser(context.Background(), &TestListUserReq)
	if err != nil {
		t.Error(err)
		t.Failed()
	}

	log.Println(users)
}

func TestGRPCCreateUser(t *testing.T) {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.Dial(addr, opts...)
	defer conn.Close()

	if err != nil {
		t.Error(err)
		t.Failed()
	}

	clt := pb.NewUserClient(conn)

	objID, err := clt.CreateUser(context.Background(), &TestCreateUserReq)

	if err != nil {
		t.Error(err)
		t.Failed()
	}

	log.Println(*objID)
}

func TestGRPCRegister(t *testing.T) {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.Dial(addrtoken, opts...)
	defer conn.Close()

	if err != nil {
		t.Error(err)
		t.Failed()
	}

	clt := pbt.NewTokenClient(conn)

	_, err = clt.RegisterJWTSecret(context.Background(), &TestRegisterJWTSecretRequest)

	if err != nil {
		t.Error(err)
		t.Failed()
	}
}

func TestIssueJWTToken(t *testing.T) {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.Dial(addrtoken, opts...)
	defer conn.Close()

	if err != nil {
		t.Error(err)
		t.Failed()
	}

	clt := pbt.NewTokenClient(conn)

	tkn, err := clt.IssueJWTToken(context.Background(), &TestIssueTokenReq)
	if err != nil {
		t.Error(err)
		t.Failed()
	}

	log.Println(tkn.Id)
	log.Println(tkn.AccessToken)
	log.Println(tkn.RefreshToken)
}

func TestValidateJWTToken(t *testing.T) {
	should := require.New(t)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.Dial(addrtoken, opts...)
	defer conn.Close()

	if err != nil {
		t.Error(err)
		t.Failed()
	}
	should.NoError(err)

	clt := pbt.NewTokenClient(conn)

	tkn, err := clt.ValidateToken(context.Background(), &pbt.ValidateTokenRequest{
		AccessToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJ1c2VyLW1hbmFnZXJfMDAwMCIsImV4cCI6MTY5NTg0NzQ3MiwiaWF0IjoxNjk1ODQwNDgyLCJpc3MiOiJVc2VyTWFuYWdlciIsIm5hbWUiOiJhcHBfYWRtaW5AdGVzdC5jb20iLCJzdWIiOiI2NTEyZTliZGM3YjVlNDliNTEzNjU3MzMiLCJ0eXAiOiJKV1QifQ.J3UvVyumnL3MdGKewYslIIr8uQ3mVnhhYqEL_WtK1rg",
		Jwtkey:      "jwt-key",
	})
	if err != nil {
		t.Error(err)
	}

	log.Println(tkn.IsValid)
	log.Println(tkn.InvalidReason)
	should.NoError(err)
	should.True(tkn.IsValid)
}

func TestRefreshJWTToken(t *testing.T) {
	should := require.New(t)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.Dial(addrtoken, opts...)
	defer conn.Close()

	if err != nil {
		t.Error(err)
		t.Failed()
	}
	should.NoError(err)

	clt := pbt.NewTokenClient(conn)

	tkn, err := clt.RefreshToken(context.Background(), &pbt.RefreshTokenRequest{
		AccessToken:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJ1c2VyLW1hbmFnZXJfMDAwMCIsImV4cCI6MTY5NTg0MTM4MiwiaWF0IjoxNjk1ODQwNDgyLCJpc3MiOiJVc2VyTWFuYWdlciIsIm5hbWUiOiJhcHBfYWRtaW5AdGVzdC5jb20iLCJzdWIiOiI2NTEyZTliZGM3YjVlNDliNTEzNjU3MzMiLCJ0eXAiOiJKV1QifQ.0G8KMzSKvesVmmBjsNR4D62UcnEoVhCSOaMFdX9iLcw",
		RefreshToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTU5MjE2MDIsImlhdCI6MTY5NTgzNTIwMiwiaXNzIjoiVXNlck1hbmFnZXIiLCJzdWIiOiI2NTEyZTliZGM3YjVlNDliNTEzNjU3MzMifQ.GkTrZ7SxNIqGTlwRl1xcRPo3ZLts1OpWtuEEgafwo44",
		Jwtkey:       "jwt-key",
	})
	if err != nil {
		t.Error(err)
	}

	log.Println(tkn.AccessToken)
	should.True(tkn.Ok)
	should.NoError(err)

}
