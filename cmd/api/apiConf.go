package api

import (
	"log"
	"win/envoice/internal/database/mongodb"
)

// type ApiConfig struct {
// 	M                 mongodb.MongoDB
// 	InfoLog, ErrorLog *log.Logger
// }

type ApiConfig struct {
	M                 mongodb.Envoicer
	InfoLog, ErrorLog *log.Logger
}
