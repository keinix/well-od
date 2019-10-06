package payment

import "testing"

func TestUsd_FromString(t *testing.T) {
	tables := []struct {
		input    string
		expected Usd
	}{
		{input: "$15.27", expected: Usd(1527)},
		{input: "6.00", expected: Usd(600)},
		{input: "0.00", expected: Usd(0)},
		{input: "-1,237.60,", expected: Usd(123760)},
	}

	for _, table := range tables {
		var u Usd
		u.FromString(table.input)
		if u != table.expected {
			t.Errorf("Input: %v -- Expected: %v Actual: %v",table.input, table.expected, u)
		}
	}
}

func TestUsd_ToString(t *testing.T) {
	u := Usd(199)
	actual := u.ToString()
	expected := "$1.99"
	if actual != expected {
		t.FailNow()
	}
}
