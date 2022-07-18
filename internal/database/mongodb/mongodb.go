package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	databaseName string = "win"
	collection          = "envoice"
)

// MongoDB
type MongoDB struct {
	Mdb               *mongo.Client
	InfoLog, ErrorLog *log.Logger
}

func Connection() (*mongo.Client, error) {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	// defer client.Disconnect(ctx)

	return client, nil
}
