package router

import (
	"net/http"
	"win/envoice/cmd/api"
)

func Router(a *api.ApiConfig) http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("/api/env", a.EnvoiceHandler)
	mux.HandleFunc("/api/env/num", a.GetEnvoiceHandler)

	return mux

}
