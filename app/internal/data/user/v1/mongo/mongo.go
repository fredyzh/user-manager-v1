package mongo

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	meta "github.com/fredyzh/user_manager/api/meta/v1/pb"
	pbt "github.com/fredyzh/user_manager/api/token/v1/pb"
	pb "github.com/fredyzh/user_manager/api/user/v1/pb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoDbUri string

type MongoDB struct {
	Host      string
	Port      string
	DefualtDb string
	UserDB    string
	TokenDB   string
	SecretsDB string
	Admin     string
	Password  string
	DBClient  *mongo.Client
	Timeout   time.Duration
}

type MongoUser struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	User      *pb.User           `bson:"user" json:"user"`
	CreatedAt primitive.DateTime `bson:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at"`
}

type MongoJwtSecret struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	JwtSecret *pbt.JWTSecrets    `bson:"jwt_secret" json:"jwt_secret"`
	CreatedAt primitive.DateTime `bson:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at"`
}

func (m *MongoDB) ConnectDB() interface{} {
	if m.Host == "localhost" || m.Host == "host.docker.internal" {
		mongoDbUri = fmt.Sprintf("mongodb://%s:%s", m.Host, m.Port)
	} else {
		// mongodb+srv://mongo_admin:****@cluster0.abh3pnp.mongodb.net/?authSource=admin
		mongoDbUri = fmt.Sprintf("mongodb+srv://%s:%s@%s/", m.Admin, m.Password, m.Host)
	}

	c, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoDbUri))
	if err != nil {
		slog.Error(err.Error())
	}

	slog.Info("connected to MongoDB")
	return c
}

func (m *MongoDB) CreateUser(ctx context.Context, usr *pb.User) (interface{}, error) {
	coll := getCollection(m, m.UserDB)
	c, cancel := context.WithTimeout(ctx, m.Timeout)
	defer cancel()

	r, err := coll.InsertOne(c, getMongoUser(usr))
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	return r.InsertedID, nil
}

func (m *MongoDB) GetUser(ctx context.Context, id string) (*pb.User, error) {
	coll := getCollection(m, m.UserDB)
	c, cancel := context.WithTimeout(ctx, m.Timeout)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var result MongoUser
	filter := bson.D{{Key: "_id", Value: objID}}
	opts := options.FindOne()
	err = coll.FindOne(c, filter, opts).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result.User, nil
}

func (m *MongoDB) ListUser(ctx context.Context) ([]*pb.User, error) {
	coll := getCollection(m, m.UserDB)
	c, cancel := context.WithTimeout(ctx, m.Timeout)
	defer cancel()

	csr, err := coll.Find(c, bson.M{})
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}
	defer csr.Close(ctx)
	var users []*pb.User = []*pb.User{}

	for csr.Next(ctx) {
		user := MongoUser{}
		if err = csr.Decode(&user); err != nil {
			slog.Error(err.Error())
			return nil, err
		}

		users = append(users, user.User)
	}

	return users, nil
}

func (m *MongoDB) GetUserByUsername(ctx context.Context, name string, scope *meta.Scope) (*pb.User, string, error) {
	coll := getCollection(m, m.UserDB)
	c, cancel := context.WithTimeout(ctx, m.Timeout)
	defer cancel()

	filter := bson.M{
		"$and": bson.A{
			bson.M{
				"user.basicauth.user_name": name,
			},
			bson.M{
				"$and": bson.A{
					bson.M{
						"user.scope.domain": scope.Domain,
					},
					bson.M{
						"user.scope.app_id": scope.AppId,
					},
				},
			},
		},
	}

	var result MongoUser
	err := coll.FindOne(c, filter).Decode(&result)
	if err != nil {
		return nil, "", err
	}

	return result.User, result.ID.Hex(), nil
}

func (m *MongoDB) SaveProfile(ctx context.Context, id string, prof *pb.Profile) (interface{}, error) {
	coll := getCollection(m, m.UserDB)
	c, cancel := context.WithTimeout(ctx, m.Timeout)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{Key: "_id", Value: objID}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "user.profile", Value: prof}}}}
	res, err := coll.UpdateOne(c, filter, update)

	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	return res.UpsertedCount, nil
}

func (m *MongoDB) SaveAddress(ctx context.Context, id string, addr *pb.Address) (interface{}, error) {
	coll := getCollection(m, m.UserDB)
	c, cancel := context.WithTimeout(ctx, m.Timeout)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{Key: "_id", Value: objID}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "user.profile.address", Value: addr}}}}
	res, err := coll.UpdateOne(c, filter, update)

	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	return res, nil
}

func (m *MongoDB) SaveJWTSecret(ctx context.Context, s *pbt.JWTSecrets) (interface{}, error) {
	coll := getCollection(m, m.SecretsDB)
	c, cancel := context.WithTimeout(ctx, m.Timeout)
	defer cancel()

	t := primitive.NewDateTimeFromTime(time.Now().AddDate(-1, 0, 0))
	r, err := coll.InsertOne(c, &MongoJwtSecret{
		ID:        primitive.NewObjectID(),
		JwtSecret: s,
		CreatedAt: t,
		UpdatedAt: t,
	})
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	return r.InsertedID, nil
}

func (m *MongoDB) GetJwtSecret(ctx context.Context, userId string, key string) (string, error) {
	coll := getCollection(m, m.SecretsDB)
	c, cancel := context.WithTimeout(ctx, m.Timeout)
	defer cancel()

	filter := bson.M{
		"$and": bson.A{
			bson.M{
				"jwt_secret.user_id": userId,
			},
			bson.M{
				"jwt_secret.key": key,
			},
		},
	}

	var result MongoJwtSecret
	err := coll.FindOne(c, filter).Decode(&result)
	if err != nil {
		return "", err
	}

	return result.JwtSecret.Value, nil
}

// common func
func getMongoUser(usr *pb.User) *MongoUser {
	t := primitive.NewDateTimeFromTime(time.Now().AddDate(-1, 0, 0))
	return &MongoUser{
		ID:        primitive.NewObjectID(),
		User:      usr,
		CreatedAt: t,
		UpdatedAt: t,
	}
}

func getCollection(m *MongoDB, coll string) *mongo.Collection {
	client := m.DBClient

	if client == nil {
		m.DBClient = m.ConnectDB().(*mongo.Client)
		client = m.DBClient
	}

	return client.Database(m.DefualtDb).Collection(coll)
}
