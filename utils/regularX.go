package utils

import (
	"regexp"
)

func Last4(cardNumber *string) string {

	re := regexp.MustCompile(`\d{12}`)
	last4 := re.ReplaceAll([]byte(*cardNumber), []byte("****-****-****"))

	return string(last4)

}

// ReplChars helper to raplece charecter of a card
// src target string 1111222233334444, replByChar to replace ****-***-***
// strRegX regular expresion take string literal `\d{3}`
func ReplChars(src *string, strRegX string, replByChar string) {

	re := regexp.MustCompile(strRegX)
	replByte := re.ReplaceAll([]byte(*src), []byte(replByChar))

	*src = string(replByte)

}
