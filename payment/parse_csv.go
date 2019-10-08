package payment

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
	"time"
)

const (
	SameDateEachMonth BillingPeriod = iota
	EveryThirtyDays   BillingPeriod = iota
)

type BillingPeriod int8

type Periodic struct {
	Name     string
	Amount   Usd
	Cycle    BillingPeriod
	LastPaid int64
}

type deduction struct {
	name   string
	amount Usd
	date   int64
}

func ParseCsv(path string) {
	r := getCsvRecords(path)
	deductions := getDeductions(r)
	log.Println(deductions)
	// 2. see if an deductions have
	// 3. if the names or payments are the same and the times match
	// a periodic cycle, but the amounts are different, take the most recent amount
}

func getCsvRecords(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("error reading csv from %v %v", path, err)
	}
	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatalf("error reading csv data: %v", err)
	}
	return records
 }

// Csv columns:
// 0: Date: MM/DD/YYYY
// 1: Amount
// 2, 3: unused
// 4: transaction name
func getDeductions(records [][]string) []deduction {
	var deductions []deduction
	for _, r := range records {
		if !strings.HasPrefix(r[1], "-") {
			continue
		}
		var amount Usd
		amount.FromString(r[1])
		d := deduction{
			name:   r[4],
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
