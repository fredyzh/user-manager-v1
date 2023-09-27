package grpctest

import (
	meta "github.com/fredyzh/user_manager/api/meta/v1/pb"
	pbj "github.com/fredyzh/user_manager/api/token/v1/pb"
	pbt "github.com/fredyzh/user_manager/api/token/v1/service"
	usr "github.com/fredyzh/user_manager/api/user/v1/pb"
	pb "github.com/fredyzh/user_manager/api/user/v1/service"
)

var (
	objId                        = "650d054d220eef74c13a17a7"
	TestGetUserReq pb.GetUserReq = pb.GetUserReq{
		Id: objId,
	}

	TestGetUserByUsernameReq pb.GetUserByUsernameReq = pb.GetUserByUsernameReq{
		Username: "app_admin@test.com",
	}

	TestResetUserPasswordReq pb.ResetUserPasswordReq = pb.ResetUserPasswordReq{
		Id:          objId,
		OldPassword: "test1234",
		NewPassword: "test2345",
	}

	TestCreateUserReq pb.CreateUserReq = pb.CreateUserReq{
		Username: "app_admin@test.com",
		Password: "test1234",
		Scope: &meta.Scope{
			Domain: "user-manager",
			AppId:  "0000",
		},
	}

	TestIssueTokenReq pbt.IssueTokenRequest = pbt.IssueTokenRequest{
		Username: "app_admin@test.com",
		Password: "test1234",
		Jwtkey:   "jwt-key",
		Scope: &meta.Scope{
			Domain: "user-manager",
			AppId:  "0000",
		},
	}

	TestListUserReq pb.ListUserReq = pb.ListUserReq{}

	TestUser usr.User = usr.User{
		BasicAuth: &usr.BasicAuth{
			Username: "app_admin@test.com",
			Password: "test1234",
		},
		Scope: &meta.Scope{
			Domain: "user-manager",
			AppId:  "0000",
		},
	}

	TestVerifyPasswordReq pb.VerifyPasswordReq = pb.VerifyPasswordReq{
		Username: "app_admin@test.com",
		Password: "test1234",
	}

	TestAddress usr.Address = usr.Address{
		Street: "somewhere rd.",
		City:   "buffalo",
		State:  "NY",
		Zip:    "14356",
	}

	TestProfile usr.Profile = usr.Profile{
		FirstName: "app1",
		LastName:  "admin",
		Email:     "app_admin@test.com",
		Phone:     "12345678901",
		Cell:      "12345678901",
		Gender:    usr.Gender_MALE,
		Address:   &TestAddress,
	}

	TestSaveProfileReq pb.SaveProfileReq = pb.SaveProfileReq{
		Id:      objId,
		Profile: &TestProfile,
	}

	TestGetProfileReq pb.GetProfileReq = pb.GetProfileReq{
		Id: objId,
	}

	TestCreateAddressReq pb.SaveAddressReq = pb.SaveAddressReq{
		Id:      objId,
		Address: &TestAddress,
	}

	TestGetAddressReq pb.GetAddressReq = pb.GetAddressReq{
		Id: objId,
	}

	TestRegisterJWTSecretRequest pbt.RegisterJWTSecretRequest = pbt.RegisterJWTSecretRequest{
		Username: "app_admin@test.com",
		Password: "test1234",
		Scope: &meta.Scope{
			Domain: "user-manager",
			AppId:  "0000",
		},
		Jwt: &pbj.JWTSecrets{
			Key:   "jwt-key",
			Value: "jwt-test-secrets",
		},
	}
)
