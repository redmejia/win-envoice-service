package api

import (
	"log"
	"win/envoice/internal/database/mongodb"
)

type ApiConfig struct {
	M                 mongodb.MongoDB
	InfoLog, ErrorLog *log.Logger
}
