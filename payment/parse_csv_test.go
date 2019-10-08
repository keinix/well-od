package payment

import "testing"

func TestGetDeductions(t *testing.T) {
	recordsWithDeductions := [][]string{
		{"09/20/2019", "-100.21", "*", "", "Big bean burrito"},
		{"09/21/2019", "2000.11", "*", "", "Direct deposit"},
		{"09/22/2019", "-14.99", "*", "", "Netflix"},
	}
	recordsWithNoDeductions := [][]string{
		{"09/20/2019", "300.00", "*", "", "Birthday"},
		{"09/21/2019", "2000.11", "*", "", "Direct deposit"},
		{"09/22/2019", "50.00", "*", "", "Rebate"},
	}
	tables := []struct {
		name string
		input [][]string
		want int // number of deductions
	}{
		{name: "records with deductions", input: recordsWithDeductions, want: 2},
		{name: "records with no deductions", input: recordsWithNoDeductions, want: 0},
	}
	for _, table := range tables {
		got := getDeductions(table.input)
		if len(got) != table.want {
			t.Errorf("%v -> got: %v want: %v", table.name, len(got), table.want)
		}
	}
}

func TestUnixTimeFromRecords(t *testing.T) {
	input := "12/25/2019"
	var want int64 = 1577232000
	got := unixTimeFromRecord(input)
	if got != want {
		t.Errorf("input %v got: %d want: %d", input, got, want)
	}
}