package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type db struct {
	*mongo.Collection
}

func New() (*db, error) {
	opts := options.Client().ApplyURI("mongodb://test:test@localhost:27017/test?ssl=false")
	err := opts.Validate()
	if err != nil {
		return nil, err
	}
	c, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err = c.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	database := c.Database("test")
	return &db{
		Collection: database.Collection("test"),
	}, nil
}
