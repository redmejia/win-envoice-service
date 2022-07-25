package api

import (
	"bytes"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"win/envoice/internal/models"
	"win/envoice/utils"
)

var errNoDocuments = errors.New("no document not in result")

type mockDatabase struct {
	Mdb               map[string][]models.TransactionData
	InfoLog, ErrorLog *log.Logger
}

func newMockDatabase(infoLog, errLog *log.Logger) *mockDatabase {
	var db = make(map[string][]models.TransactionData)

	var envData = models.TransactionData{
		EnvoiceUUID:  "e25711d7-7166-4141-bb3a-23fa1d3ac848",
		TxAccepted:   true,
		MessageState: "Transaction Accecpted",
		Date:         "2022-07-24 18:52:49.057280468",
		Transaction: models.Transaction{
			Product: models.ProductSpecification{
				ProductName:  "Cars",
				ProductPrice: 100,
			},
			TxAmount:     100,
			TxDate:       "2022-07-24 18:52:49.057280468",
			TxCardNumber: "****-****-****-2369",
			TxCardCv:     "***",
			BillingInfo: models.BillingInfo{
				FullName: "Elon Musk",
				Address:  "108 main st",
				City:     "Fremont",
				State:    "CA",
				Zip:      "123456",
			},
		},
	}

	var document []models.TransactionData
	document = append(document, envData)
	db["win"] = document

	return &mockDatabase{
		Mdb:      db,
		InfoLog:  infoLog,
		ErrorLog: errLog,
	}

}

func insertDoc(db mockDatabase, doc models.TransactionData) {
	var document []models.TransactionData

	document = append(document, doc)
	db.Mdb["win"] = document

}

func findDoc(db mockDatabase, envoiceUUID string) (*models.TransactionData, error) {
	var doc []models.TransactionData

	for _, v := range db.Mdb["win"] {
		if v.EnvoiceUUID == envoiceUUID {
			doc = append(doc, v)
		} else {
			return nil, errNoDocuments
		}
	}

	return &doc[0], nil

}

func (m *mockDatabase) CreateEnvoiceRecord(envObj models.TransactionData) {
	insertDoc(*m, envObj)
}

func (m *mockDatabase) GetEnvoiceByUUID(w http.ResponseWriter, envoiceUUID string) {

	doc, err := findDoc(*m, envoiceUUID)
	if err != nil {
		if errors.Is(err, errNoDocuments) {
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

	err = utils.WriteJSON(w, http.StatusOK, &doc)
	if err != nil {
		m.ErrorLog.Fatal(err)
	}

}

func TestGetEnvoiceHandlerSuccess(t *testing.T) {

	req, err := http.NewRequest("GET", "/api/env/num?envo-uuid=e25711d7-7166-4141-bb3a-23fa1d3ac848", nil)
	if err != nil {
		t.Error(err)
	}

	recorder := httptest.NewRecorder()

	handler := http.HandlerFunc(app.GetEnvoiceHandler)

	handler.ServeHTTP(recorder, req)
	want := `{"id":"000000000000000000000000","envoice_uuid":"e25711d7-7166-4141-bb3a-23fa1d3ac848","tx_accepted":true,"message_state":"Transaction Accecpted","date":"2022-07-24 18:52:49.057280468","transaction":{"product":{"product_name":"Cars","product_price":100},"tx_amount":100,"tx_date":"2022-07-24 18:52:49.057280468","tx_card_number":"****-****-****-2369","tx_card_cv":"***","billing_info":{"full_name":"Elon Musk","address":"108 main st","city":"Fremont","state":"CA","zip":"123456"}}}`

	if recorder.Code != http.StatusOK && recorder.Body.String() != want {
		t.Errorf("not found want %s but %s", want, recorder.Body.String())
	}

}

func TestGetEnvoiceHandlerBadUUID(t *testing.T) {

	req, err := http.NewRequest("GET", "/api/env/num?envo-uuid=b3e8a34d", nil)
	if err != nil {
		t.Error(err)
	}

	recorder := httptest.NewRecorder()

	handler := http.HandlerFunc(app.GetEnvoiceHandler)

	handler.ServeHTTP(recorder, req)

	wantError := `{"is_error":true,"err_msg":"envoice uuid was not found"}`

	if recorder.Code == http.StatusNotFound && recorder.Body.String() == wantError {
		t.Log("match expected error")
	}

}

func TestEnvoiceHandler(t *testing.T) {

	envoice := `{"tx_accepted":true,"message_state":"Transaction Accecpted","date":"2022-07-24 18:52:49.057280468","transaction":{"product":{"product_name":"Cars","product_price":100},"tx_amount":100,"tx_date":"2022-07-24 18:52:49.057280468","tx_card_number":"1111222233332369","tx_card_cv":"103","billing_info":{"full_name":"Elon Musk","address":"108 main st","city":"Fremont","state":"CA","zip":"123456"}}}`

	req, err := http.NewRequest("POST", "/api/env", bytes.NewBuffer([]byte(envoice)))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	handler := http.HandlerFunc(app.EnvoiceHandler)

	handler.ServeHTTP(recorder, req)

	if recorder.Code == http.StatusCreated {
		t.Log(recorder.Body.String())
	}
}
