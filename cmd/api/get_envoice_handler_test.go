package api

// func TestEnvoiceHandler(t *testing.T) {

// 	envData := models.TransactionData{
// 		TxAccepted:   true,
// 		MessageState: "Transaction Accecpted",
// 		Date:         time.Now(),
// 		Transaction: models.Transaction{
// 			Product: models.ProductSpecification{
// 				ProductName:  "Cars",
// 				ProductPrice: 100,
// 			},
// 			TxAmount:     100,
// 			TxDate:       time.Now(),
// 			TxCardNumber: "1111222233332369",
// 			TxCardCv:     "103",
// 			BillingInfo: models.BillingInfo{
// 				FullName: "Elon Musk",
// 				Address:  "108 main st",
// 				City:     "Fremont",
// 				State:    "CA",
// 				Zip:      "123456",
// 			},
// 		},
// 	}

// 	data, err := json.Marshal(envData)
// 	if err != nil {
// 		t.Log(err)
// 		return
// 	}

// 	req, err := http.NewRequest("POST", "/api/env", bytes.NewBuffer(data))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	req.Header.Set("Content-Type", "application/json")

// 	recorder := httptest.NewRecorder()

// 	handler := http.HandlerFunc(app.EnvoiceHandler)

// 	handler.ServeHTTP(recorder, req)

// 	if recorder.Code == http.StatusCreated {
// 		t.Log("ok created")
// 	}
// }

// func TestGetEnvoiceHandlerSuccess(t *testing.T) {

// 	req, err := http.NewRequest("GET", "/api/env/num?envo-uuid=b3e8a34d-f3d4-4a9d-b2d3-0c024445f8f1", nil)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	recorder := httptest.NewRecorder()

// 	handler := http.HandlerFunc(app.GetEnvoiceHandler)

// 	handler.ServeHTTP(recorder, req)

// 	want := `{"_id":"62d8d6bd86bb70b0ba642ed6","date":"2022-07-20T21:31:57.396-07:00","envoice_uuid":"b3e8a34d-f3d4-4a9d-b2d3-0c024445f8f1","message_state":"Transanction Accepted","transaction":{"billing_info":{"address":"1412 main ave","city":"Oakland","full_name":"Elon Musk","state":"cal","zip":"535353"},"product":{"product_name":"cars","product_price":100},"tx_amount":100,"tx_card_cv":"103","tx_card_number":"****-****-****2369","tx_date":"2022-07-20T21:31:57.396-07:00"},"tx_accepted":true}`

// 	if recorder.Code != http.StatusOK && recorder.Body.String() != want {
// 		t.Errorf("recived %s \n want %s", recorder.Body.String(), want)
// 	}

// }

// func TestGetEnvoiceHandlerBadUUID(t *testing.T) {

// 	req, err := http.NewRequest("GET", "/api/env/num?envo-uuid=b3e8a34d", nil)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	recorder := httptest.NewRecorder()

// 	handler := http.HandlerFunc(app.GetEnvoiceHandler)

// 	handler.ServeHTTP(recorder, req)

// 	wantError := `{"is_error":true,"err_msg":"envoice uuid was not found"}`

// 	if recorder.Code == http.StatusNotFound && recorder.Body.String() == wantError {
// 		t.Log("match expected error")
// 	}

// }
