package api

import (
	"encoding/json"
	"net/http"
	"win/envoice/internal/models"
	"win/envoice/utils"

	"github.com/google/uuid"
)

func (a *ApiConfig) EnvoiceHandler(w http.ResponseWriter, r *http.Request) {
	var transaction models.TransactionData

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&transaction)
	if err != nil {
		a.ErrorLog.Fatal(err)
	}

	envUUID := uuid.NewString()
	transaction.EnvoiceUUID = envUUID

	last4 := utils.Last4(&transaction.Transaction.TxCardNumber)

	transaction.Transaction.TxCardNumber = last4

	a.M.CreateEnvoiceRecord(transaction)

	// sent back the envUUID this will be save on the database of bussiness logic
	// retrive when GET reques is made to win envoice/id

	envoice := models.EnvoiceInfo{Recived: true, RecordCreated: true, EnvoUUID: envUUID}

	a.InfoLog.Println(envoice)

	envByte, err := json.Marshal(envoice)
	if err != nil {
		a.ErrorLog.Fatal(err)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write(envByte)

}
