package mongodb

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"
	"win/envoice/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	databaseName  string = "win"
	envCollection        = "envoice" // envoice collection
)

// MongoDB
type MongoDB struct {
	Mdb               *mongo.Client
	InfoLog, ErrorLog *log.Logger
}

type Envoicer interface {
	CreateEnvoiceRecord(envObj models.TransactionData)
	GetEnvoiceByUUID(w http.ResponseWriter, envoiceUUID string)
}

func Connection() (*mongo.Client, error) {

	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("URI")))

	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	// defer client.Disconnect(ctx)

	return client, nil
}
