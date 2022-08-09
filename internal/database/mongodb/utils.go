package mongodb

import (
	"context"
	"errors"
	"net/http"
	"time"
	"win/envoice/internal/models"
	"win/envoice/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

	singleResult := envoiceColl.FindOne(ctx, bson.M{"envoice_uuid": envoiceUUID}, nil)
	err := singleResult.Decode(&env)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			var badUUID = struct {
				IsError bool   `json:"is_error"`
				ErrMsg  string `json:"err_msg"`
			}{
				IsError: true,
				ErrMsg:  "envoice uuid was not found",
			}

			err = utils.WriteJSON(w, http.StatusNotFound, badUUID)
			if err != nil {
				m.ErrorLog.Fatal(err)
			}
		}
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, &env)
	if err != nil {
		m.ErrorLog.Fatal(err)
		return
	}
}

func (m *MongoDB) GetAllByCompanyUid(w http.ResponseWriter, companyUID string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	envoiceColl := m.Mdb.Database(databaseName).Collection(envCollection)

	// find all by company uid enter dot notation
	cursor, err := envoiceColl.Find(ctx, bson.M{"transaction.company_uid": companyUID})
	if err != nil {
		m.ErrorLog.Fatal("error here ", err)
	}

	var result []bson.M

	err = cursor.All(ctx, &result)
	if err != nil {
		m.ErrorLog.Fatal("error also here", err)
	}

	m.InfoLog.Println(result)

	err = utils.WriteJSON(w, http.StatusFound, &result)
	if err != nil {
		m.ErrorLog.Fatal(err)
	}

}
