package api

import (
	"log"
	"os"
	"testing"
	"win/envoice/internal/database/mongodb"
)

var app ApiConfig

func TestMain(m *testing.M) {
	client, _ := mongodb.Connection()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime)

	app.M = mongodb.MongoDB{
		Mdb:      client,
		InfoLog:  infoLog,
		ErrorLog: errLog,
	}

	app.InfoLog = infoLog
	app.ErrorLog = errLog

	code := m.Run()

	os.Exit(code)
}
