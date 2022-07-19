package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"win/envoice/internal/models"

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
	a.M.CreateEnvoiceRecord(txEnvodata)

	fmt.Println("data recived ", txEnvodata)
}
