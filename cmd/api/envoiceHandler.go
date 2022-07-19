package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"win/envoice/internal/models"
	"win/envoice/utils"

	"github.com/google/uuid"
)

func (a *ApiConfig) EnvoiceHandler(w http.ResponseWriter, r *http.Request) {
	var txEnvodata models.TransactionData

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&txEnvodata)
	if err != nil {
		a.ErrorLog.Fatal(err)
	}

	envUUID := uuid.NewString()
	txEnvodata.EnvoiceUUID = envUUID

	last4 := utils.Last4(&txEnvodata.Transaction.TxCardNumber)

	txEnvodata.Transaction.TxCardNumber = last4

	a.M.CreateEnvoiceRecord(txEnvodata)

	fmt.Println("data recived ", txEnvodata)
}
