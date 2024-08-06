package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewConnection(c context.Context, uri string) (*mongo.Client, error) {
	opts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(c, opts)
	if err != nil {
		return nil, err
	}
	err = client.Ping(c, nil)
	if err != nil {
		return nil, err
	}
	return client, nil
}
