package api

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type ApiConfig struct {
	Mdb               *mongo.Client
	InfoLog, ErrorLog *log.Logger
}
