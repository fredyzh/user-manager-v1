package grpctest

import (
	"context"
	"log"
	"testing"

	pbt "github.com/fredyzh/user_manager/api/token/v1/service"
	pb "github.com/fredyzh/user_manager/api/user/v1/service"
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
	log.Println(tkn)
}
