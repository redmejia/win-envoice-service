package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ProductSpecification
type ProductSpecification struct {
	ProductName  string `bson:"product_name,omitempty" json:"product_name"`
	ProductPrice int    `bson:"product_price,omitempty" json:"product_price"`
}

// BillingInfo
type BillingInfo struct {
	FullName string `bson:"full_name,omitempty" json:"full_name"`
	Address  string `bson:"address,omitempty" json:"address"`
	City     string `bson:"city,omitempty" json:"city"`
	State    string `bson:"state,omitempty" json:"state"`
	Zip      string `bson:"zip,omitempty" json:"zip"`
}

// Transaction
type Transaction struct {
	CompanyUID   string               `bson:"company_uid,omitempty" json:"company_uid"`
	Product      ProductSpecification `bson:"product,omitempty" json:"product"` // Business product specification
	TxAmount     int                  `bson:"tx_amount,omitempty" json:"tx_amount"`
	TxDate       string               `bson:"tx_date,omitempty" json:"tx_date"`
	TxCardNumber string               `bson:"tx_card_number,omitempty" json:"tx_card_number"`
	TxCardCv     string               `bson:"tx_card_cv,omitempty" json:"tx_card_cv"`
	BillingInfo  BillingInfo          `bson:"billing_info,omitempty" json:"billing_info"`
}

// TransactionStatus
type TransactionStatus struct {
	Proceed        bool   `bson:"proceed,omitempty" json:"proceed"`
	TxAmountIntent int    `bson:"tx_amount_intent,omitempty" json:"tx_amount_intent"`
	TxStatusCode   int    `bson:"tx_status_code,omitempty" json:"tx_status_code"`
	TxMessage      string `bson:"tx_message,omitempty" json:"tx_message"`
}

// TransactioData
type TransactionData struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	EnvoiceUUID  string             `bson:"envoice_uuid,omitempty" json:"envoice_uuid"`
	TxAccepted   bool               `bson:"tx_accepted,omitempty" json:"tx_accepted"`
	MessageState string             `bson:"message_state,omitempty" json:"message_state"`
	Date         string             `bson:"date,omitempty" json:"date"`
	Transaction  Transaction        `bson:"transaction,omitempty" json:"transaction"`
}

// EnvoiceInfo
type EnvoiceInfo struct {
	Recived       bool   `json:"recived"`        // request was recived
	RecordCreated bool   `json:"record_created"` // new record was created no metter accepted or declined
	EnvoUUID      string `json:"envo_uuid"`
}
