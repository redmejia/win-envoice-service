package mongodb

import (
	"context"
	"log"
	"time"
	"win/envoice/internal/models"
)

func (m *MongoDB) CreateEnvoiceRecord(envObj models.TransactionData) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	defer m.Mdb.Disconnect(ctx)

	database := m.Mdb.Database(databaseName)
	envoiceColl := database.Collection(envCollection)

	result, err := envoiceColl.InsertOne(ctx, envObj)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("id ", result.InsertedID)
}
