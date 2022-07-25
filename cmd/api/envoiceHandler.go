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

	if r.Method == http.MethodPost {

		var transaction models.TransactionData

		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&transaction)
		if err != nil {
			a.ErrorLog.Fatal(err)
		}

		envUUID := uuid.NewString()
		transaction.EnvoiceUUID = envUUID

		utils.ReplChars(&transaction.Transaction.TxCardNumber, `\d{12}`, "****-****-****-")
		utils.ReplChars(&transaction.Transaction.TxCardCv, `\d{3}`, "***")

		a.M.CreateEnvoiceRecord(transaction)

		// sent back the envUUID this will be save on the database of bussiness logic
		// retrive when GET request is made to win envoice/id

		envoice := models.EnvoiceInfo{Recived: true, RecordCreated: true, EnvoUUID: envUUID}

		// a.InfoLog.Println(envoice)

		err = utils.WriteJSON(w, http.StatusCreated, &envoice)

		if err != nil {
			a.ErrorLog.Fatal(err)
		}
	} else {
		var methodNotImplementd = struct {
			Error        bool   `json:"error"`
			ErrorMessage string `json:"error_message"`
		}{
			Error:        true,
			ErrorMessage: fmt.Sprintf(`%s is not supported`, r.Method),
		}
		err := utils.WriteJSON(w, http.StatusNotImplemented, &methodNotImplementd)
		if err != nil {
			a.ErrorLog.Fatal(err)
		}
	}

}

func (a *ApiConfig) GetEnvoiceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// http: //localhost:8089/api/env/num?envo-uuid=1233sdh-313030-12312-313
		envUUID := r.URL.Query().Get("envo-uuid")

		a.M.GetEnvoiceByUUID(w, envUUID)
	} else {
		var methodNotImplementd = struct {
			Error        bool   `json:"error"`
			ErrorMessage string `json:"error_message"`
		}{
			Error:        true,
			ErrorMessage: fmt.Sprintf(`%s is not supported`, r.Method),
		}
		err := utils.WriteJSON(w, http.StatusNotImplemented, &methodNotImplementd)
		if err != nil {
			a.ErrorLog.Fatal(err)
		}
	}

}
