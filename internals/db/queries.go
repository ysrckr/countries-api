package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *service) QueryAll(ctx context.Context, collection string) (*mongo.Cursor, error) {
	coll := s.db.Database(s.dbName).Collection(collection)
	cursor, err := coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	return cursor, nil
}
