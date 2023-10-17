package v1

import (
	"context"
	"log"
	"log/slog"
	"testing"
	"time"

	meta "github.com/fredyzh/user_manager/api/meta/v1/pb"
	pbt "github.com/fredyzh/user_manager/api/token/v1/pb"
	pb "github.com/fredyzh/user_manager/api/user/v1/pb"
	mongodb "github.com/fredyzh/user_manager/app/internal/data/user/v1/mongo"
	. "github.com/smartystreets/goconvey/convey"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mon = &mongodb.MongoDB{
	Host:      "cluster0.abh3pnp.mongodb.net",
	Port:      "27017",
	DefualtDb: "UserManagerV3",
	UserDB:    "user",
	TokenDB:   "token",
	SecretsDB: "secrets",
	Admin:     "mongo_admin",
	Password:  "GodIsLove!",
	Timeout:   10 * time.Second,
}

var scope = &meta.Scope{
	Domain: "user-manager",
	AppId:  "0000",
}

var usrBasicAuth = &pb.BasicAuth{
	Username: "app_admin@test.com",
	Password: "test1234",
}

var user = &pb.User{
	BasicAuth: usrBasicAuth,
	Scope:     scope,
}

// var role = &models.Role{
// 	RoleName: "app_amdin",
// 	RoleType: models.RoleType_APP_ADMIN,
// }

// var thirdPartySecret = &models.ThirdPartySecret{
// 	KeyName:  "JwtSecret4Test1",
// 	KeyValue: "jwt-Secret_1",
// }

var profile = &pb.Profile{
	FirstName: "app1",
	LastName:  "admin",
	Email:     "app_admin@test.com",
	Phone:     "12345678901",
	Cell:      "12345678901",
	Gender:    pb.Gender_MALE,
	Address:   address,
}

var address = &pb.Address{
	Street: "somewhere rd.",
	City:   "buffalo",
	State:  "NY",
	Zip:    "14356",
}

var objId = "6512e9bdc7b5e49b51365733"

// var t = primitive.NewDateTimeFromTime(time.Now().AddDate(-1, 0, 0))
var jwtSecrets = &pbt.JWTSecrets{
	UserId: objId,
	Key:    "jwt-key",
	Value:  "jwt-test-secrets",
}

func TestMongoDBConnectDB(t *testing.T) {
	var srv DataRepo = mon
	Convey("Test DB Access", t, func() {
		clt := srv.ConnectDB().(*mongo.Client)
		So(clt, ShouldNotBeNil)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err := clt.Ping(ctx, readpref.Primary())
		if err != nil {
			slog.Error(err.Error())
		}
		So(err, ShouldBeNil)
	})
}

func TestCreateUser(t *testing.T) {

	Convey("Test create user", t, func() {
		var srv DataRepo = mon

		usr := pb.User{
			BasicAuth: usrBasicAuth,
			// Scope:     scope,
			// Profile:   profile,
		}
		r, err := srv.CreateUser(context.Background(), &usr)
		log.Println(r)
		So(err, ShouldBeNil)
	})
}

func TestGetUser(t *testing.T) {

	Convey("Test get user", t, func() {
		var srv DataRepo = mon

		r, err := srv.GetUser(context.Background(), objId)
		if err != nil {
			slog.Error(err.Error())
		}
		log.Println(r)
		So(err, ShouldBeNil)
	})
}

func TestListUsers(t *testing.T) {
	Convey("Test GetUsers", t, func() {
		var srv DataRepo = mon

		usrs, err := srv.ListUser(context.Background())
		if err != nil {
			slog.Error(err.Error())
		}
		log.Println(usrs)
		So(len(usrs), ShouldBeGreaterThan, 0)
	})
}

func TestGetUserByname(t *testing.T) {

	Convey("Test user by name", t, func() {
		var srv DataRepo = mon

		r, _, err := srv.GetUserByUsername(context.Background(), "app_admin@test.com", scope)
		log.Println(r)
		So(err, ShouldBeNil)
	})
}

func TestSaveProfile(t *testing.T) {

	Convey("TestSaveProfile", t, func() {
		var srv DataRepo = mon
		r, err := srv.SaveProfile(context.Background(), "650da26cd064b32869013c78", profile)
		log.Println(r)
		So(err, ShouldBeNil)
	})
}

func TestSaveJwtSecrets(t *testing.T) {

	Convey("TestSaveJwtSecrets", t, func() {
		var srv DataRepo = mon

		r, err := srv.SaveJWTSecret(context.Background(), jwtSecrets)
		log.Println(r)
		So(err, ShouldBeNil)
	})
}

func TestGetJwtSecrets(t *testing.T) {

	Convey("TestGetJwtSecrets", t, func() {
		var srv DataRepo = mon

		r, err := srv.GetJwtSecret(context.Background(), "6512e9bdc7b5e49b51365733", "jwt-key")
		log.Println(r)
		So(err, ShouldBeNil)
	})
}
