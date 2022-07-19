package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"win/envoice/internal/models"
)

func (a *ApiConfig) EnvoiceHandler(w http.ResponseWriter, r *http.Request) {
	var txEnvodata models.TransactionData

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&txEnvodata)
	if err != nil {
		a.ErrorLog.Fatal(err)
	}

	a.M.CreateEnvoiceRecord(txEnvodata)

	fmt.Println("data recived ", txEnvodata)
}
