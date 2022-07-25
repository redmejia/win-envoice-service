package api

import (
	"log"
	"os"
	"testing"
)

var app ApiConfig

func TestMain(m *testing.M) {

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime)

	app.M = newMockDatabase(infoLog, errLog)

	app.InfoLog = infoLog
	app.ErrorLog = errLog

	code := m.Run()

	os.Exit(code)
}
