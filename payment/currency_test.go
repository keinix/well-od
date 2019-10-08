package payment

import "testing"

func TestUsd_FromString(t *testing.T) {
	tables := []struct {
		name     string
		input    string
		expected Usd
	}{
		{name: "leading $", input: "$15.27", expected: Usd(1527)},
		{name: "only numbers", input: "6.00", expected: Usd(600)},
		{name: "zero amount", input: "0.00", expected: Usd(0)},
		{name: "negative number", input: "-1,237.60,", expected: Usd(123760)},
	}

	for _, table := range tables {
		var u Usd
		u.FromString(table.input)
		if u != table.expected {
			t.Errorf("%v -> Input: %v -- Expected: %v Actual: %v", table.name, table.input, table.expected, u)
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
