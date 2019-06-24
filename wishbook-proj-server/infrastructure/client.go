package infrastructure

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newClient(ctx context.Context) (*mongo.Client, error) {

	// Create client
	client, err := mongo.NewClient(options.Client().ApplyURI(dbPath))
	if err != nil {
		return nil, err
	}

	// Create connect
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}
