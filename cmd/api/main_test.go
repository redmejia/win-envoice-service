package api

import (
	"os"
	"testing"
	"win/envoice/internal/database/mongodb"
)

var app ApiConfig

func TestMain(m *testing.M) {
	client, _ := mongodb.Connection()
	app.M = mongodb.MongoDB{
		Mdb: client,
	}

	code := m.Run()

	os.Exit(code)
}
