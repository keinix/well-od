package payment

import (
	"log"
	"time"
)

const (
	SameDateEachMonth Cycle = iota
	EveryThirtyDays   Cycle = iota
)

type Cycle int8

type Periodic struct {
	Name     string
	Amount   Usd
	Cycle    Cycle
	LastPaid int64
}

type deduction struct {
	name   string
	amount Usd
	date   int64
}

// Csv columns:
// 0: Date: MM/DD/YYYY
// 2: Amount
// 3, 4: unused
// 5: transaction name
func ParseCsv(records [][]string) {
	//deductions := getDeductions(records)
	// 2. see if an deductions have
	// 3. if the names or payments are the same and the times match
	// a periodic cycle, but the amounts are different, take the most recent amount
}

func getDeductions(records [][]string) []deduction {
	var deductions []deduction
	for _, r := range records {
		if records[0][0] != "-" {
			continue
		}
		var amount Usd
		amount.FromString(r[2])
		d := deduction{
			name:   r[5],
			amount: amount,
			date:   unixTimeFromRecord(r[0]),
		}
		deductions = append(deductions, d)
	}
	return deductions
}

func unixTimeFromRecord(s string) int64 {
	format := "01/02/2006"
	t, err := time.Parse(format, s)
	if err != nil {
		log.Fatalf("date in csv (%v) is not in MM/DD/YYYY format", s)
	}
	return t.Unix()
}
