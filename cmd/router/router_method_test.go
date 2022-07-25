package router

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"win/envoice/cmd/api"
)

var app api.ApiConfig

func TestRouterPost(t *testing.T) {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime)

	app.InfoLog = infoLog
	app.ErrorLog = errLog

	app.M = nil

	fakeData := []byte(`"test": true`)

	// none POST request fail
	req, err := http.NewRequest("PATCH", "/api/env", bytes.NewBuffer(fakeData))
	if err != nil {
		t.Fatal(err)
	}
	defer req.Body.Close()

	recoder := httptest.NewRecorder()

	handler := http.HandlerFunc(app.EnvoiceHandler)

	handler.ServeHTTP(recoder, req)

	wantError := `{"error":true,"error_message":"PATCH is not supported"}`

	if recoder.Code == http.StatusNotImplemented && recoder.Body.String() == wantError {
		t.Logf("expected error %s not implemented", req.Method)
	}
}

func TestRouterGET(t *testing.T) {

	fakeData := []byte(`"test": true`)

	// none GET request fail POST marshal error
	req, err := http.NewRequest("PUT", "/api/env/num?envo-uuid=fake", bytes.NewBuffer(fakeData))
	if err != nil {
		t.Fatal(err)
		return
	}
	defer req.Body.Close()

	recoder := httptest.NewRecorder()

	handler := http.HandlerFunc(app.GetEnvoiceHandler)

	handler.ServeHTTP(recoder, req)

	wantError := `{"error":true,"error_message":"PUT is not supported"}`

	if recoder.Code == http.StatusNotImplemented && recoder.Body.String() == wantError {
		t.Logf("expected error %s  not implemented", req.Method)
	}
}
