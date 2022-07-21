package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ProductSpecification
type ProductSpecification struct {
	ProductName  string `bson:"product_name" json:"product_name"`
	ProductPrice int    `bson:"product_price" json:"product_price"`
}

// BillingInfo
type BillingInfo struct {
	FullName string `bson:"full_name" json:"full_name"`
	Address  string `bson:"address" json:"address"`
	City     string `bson:"city" json:"city"`
	State    string `bson:"state" json:"state"`
	Zip      string `bson:"zip" json:"zip"`
}

// Transaction
type Transaction struct {
	Product      ProductSpecification `bson:"product" json:"product"` // Business product specification
	TxAmount     int                  `bson:"tx_amount" json:"tx_amount"`
	TxDate       time.Time            `bson:"tx_date" json:"tx_date"`
	TxCardNumber string               `bson:"tx_card_number" json:"tx_card_number"`
	TxCardCv     string               `bson:"tx_card_cv" json:"tx_card_cv"`
	BillingInfo  BillingInfo          `bson:"billing_info" json:"billing_info"`
}

// TransactionStatus
type TransactionStatus struct {
	Proceed        bool   `bson:"proceed" json:"proceed"`
	TxAmountIntent int    `bson:"tx_amount_intent" "json:"tx_amount_intent"`
	TxStatusCode   int    `bson:"tx_status_code" json:"tx_status_code"`
	TxMessage      string `bson:"tx_message" json:"tx_message"`
}

// TransactioData
type TransactionData struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	EnvoiceUUID  string             `bson:"envoice_uuid" json:"envoice_uuid"`
	TxAccepted   bool               `bson:"tx_accepted" json:"tx_accepted"`
	MessageState string             `bson:"message_state" json:"message_state"`
	Date         time.Time          `bson:"date" json:"date"`
	Transaction  Transaction        `bson:"transaction" json:"transaction"`
}

// EnvoiceInfo
type EnvoiceInfo struct {
	Recived       bool   `json:"recived"`        // request was recived
	RecordCreated bool   `json:"record_created"` // new record was created no metter accepted or declined
	EnvoUUID      string `json:"envo_uuid"`
}
