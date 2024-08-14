package database

import (
	"context"
	"os"

	"github.com/sv-tools/mongoifc"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoClient(c context.Context) (mongoifc.Client, error) {
	mongoDBUri := os.Getenv("MONGODB_URI")
	opts := options.Client().ApplyURI(mongoDBUri)
	client, err := mongo.Connect(c, opts)
	if err != nil {
		return nil, err
	}
	return mongoifc.WrapClient(client), nil
}
