package mongodb

import (
	"context"
	"net/http"
	"time"
	"win/envoice/internal/models"
	"win/envoice/utils"

	"go.mongodb.org/mongo-driver/bson"
)

func (m *MongoDB) CreateEnvoiceRecord(envObj models.TransactionData) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// defer m.Mdb.Disconnect(ctx)

	envoiceColl := m.Mdb.Database(databaseName).Collection(envCollection)

	_, err := envoiceColl.InsertOne(ctx, envObj)
	if err != nil {
		m.ErrorLog.Fatal(err)
		return
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

	err = utils.WriteJSON(w, http.StatusOK, &env)
	if err != nil {
		m.ErrorLog.Fatal(err)
	}
}
