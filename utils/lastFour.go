package utils

import (
	"regexp"
)

func Last4(cardNumber *string) string {

	re := regexp.MustCompile(`\d{12}`)
	last4 := re.ReplaceAll([]byte(*cardNumber), []byte("****-****-****"))

	return string(last4)

}
