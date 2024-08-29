package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *service) QueryAll(ctx context.Context, collection string, filter interface{}, options *options.FindOptions) (*mongo.Cursor, error) {
	coll := s.db.Database(s.dbName).Collection(collection)
	var cursor *mongo.Cursor
	var err error

	cursor, err = coll.Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}

	return cursor, nil
}
