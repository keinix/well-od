package payment

import (
	"log"
	"regexp"
	"strconv"
)

// U.S. dollars in cents
type Usd uint64

// takes a strings such as $23.54 are parses it to the number of cents
// inputs must contain a double digit cent value
func (u *Usd) FromString(s string) {
	numRegex, err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Fatalf("error parsing ammount field: %v", err)
	}
	digits := numRegex.ReplaceAllString(s, "")
	centsString := digits[len(digits)-2:]
	dollarsString := digits[:len(digits)-2]
	cents, err := strconv.Atoi(centsString)
	if err != nil {
		log.Fatal("error parsing cents to int")
	}
	dollars, err := strconv.Atoi(dollarsString)
	if err != nil {
		log.Fatal("error parsing dollars to int")
	}
	*u = Usd((dollars * 100) + cents)
}

func (u Usd) ToString() string {
	moneyString := strconv.Itoa(int(u))
	return "$" + moneyString[:len(moneyString)-2] + "." + moneyString[len(moneyString)-2:]
}
