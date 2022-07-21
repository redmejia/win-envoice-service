package mongodb

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
	"win/envoice/internal/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (m *MongoDB) CreateEnvoiceRecord(envObj models.TransactionData) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// defer m.Mdb.Disconnect(ctx)

	database := m.Mdb.Database(databaseName)
	envoiceColl := database.Collection(envCollection)

	_, err := envoiceColl.InsertOne(ctx, envObj)
	if err != nil {
		log.Fatal(err)
	}

}

func (m *MongoDB) GetEnvoiceByUUID(w http.ResponseWriter, envoiceUUID string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	envoiceColl := m.Mdb.Database(databaseName).Collection(envCollection)

	var env bson.M

	resutl := envoiceColl.FindOne(ctx, bson.M{"envoice_uuid": envoiceUUID}, nil)
	err := resutl.Decode(&env)
	if err != nil {
		m.ErrorLog.Fatal(err)
		return
	}
	resutl.DecodeBytes()

	b, err := json.Marshal(&env)
	if err != nil {
		m.ErrorLog.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)

}
