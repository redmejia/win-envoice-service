package main

import (
	"log"
	"net/http"
	"os"
	"win/envoice/cmd/api"
	"win/envoice/cmd/router"
	"win/envoice/internal/database/mongodb"
)

func main() {
	client, err := mongodb.Connection()
	if err != nil {
		log.Fatal(err)
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime)

	api := api.ApiConfig{
		M:        &mongodb.MongoDB{Mdb: client, InfoLog: infoLog, ErrorLog: errLog},
		InfoLog:  infoLog,
		ErrorLog: errLog,
	}

	log.Println("http://localhost:8089")
	log.Fatal(http.ListenAndServe(":8089", router.Router(&api)))

}
