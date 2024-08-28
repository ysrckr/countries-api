package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ysrckr/countries-api/internals/conf"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	config = conf.Config
	envs   = config.Envs
	DB     = New(envs["DB_USER"], envs["DB_PASS"], envs["DB_HOST"], envs["DB_PORT"], envs["DB_NAME"])
)

type Service interface {
	Health() map[string]string
	Close(context.Context) error
	QueryAll(string) *mongo.Cursor
}

type service struct {
	db     *mongo.Client
	dbName string
}

func New(username, password, host, port, dbName string) Service {
	URI := fmt.Sprintf("mongodb://%s:%s@%s:%s", username, password, host, port)

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(URI))
	if err != nil {
		log.Fatal(err)
	}
	return &service{
		db:     client,
		dbName: dbName,
	}
}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.db.Ping(ctx, nil)
	if err != nil {
		log.Fatalf(fmt.Sprintf("db down: %v", err))
	}

	return map[string]string{
		"message": "It's healthy",
	}
}

func (s *service) Close(ctx context.Context) error {
	if err := s.db.Disconnect(ctx); err != nil {
		return err
	}

	return nil
}

func (s *service) QueryAll(collection string) *mongo.Cursor {
	timeout, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	coll := s.db.Database(s.dbName).Collection(collection)
	cursor, err := coll.Find(timeout, bson.D{})
	if err != nil {
		log.Println(err)
	}

	return cursor
}
