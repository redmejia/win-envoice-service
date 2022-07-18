package main

import (
	"log"
	"net/http"
	"win/envoice/cmd/api"
	"win/envoice/cmd/router"
)

func main() {

	api := api.ApiConfig{
		Mdb:      nil,
		InfoLog:  nil,
		ErrorLog: nil,
	}

	log.Println("http://localhost:8089")
	log.Fatal(http.ListenAndServe(":8089", router.Router(&api)))
}
