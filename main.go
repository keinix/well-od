package main

import (
	"log"
	"well-od/payment"
)

func main() {
	testString := ""
	var money payment.Usd
	money.FromString(testString)
	log.Printf("result: %v expected: %v", money.ToString(), testString )
}


