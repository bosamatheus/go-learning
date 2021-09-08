package number2word

import (
	"testing"
)

func TestConvert(t *testing.T) {
	tables := []struct {
		number   int
		expected string
	}{
		{0, "zero"},
		{11, "onze"},
		{20, "vinte"},
		{34, "trinta e quatro"},
		{89, "oitenta e nove"},
		{100, "cem"},
	}

	for idx, table := range tables {
		result := Convert(table.number)
		if result != table.expected {
			t.Errorf("%d- Test failed, got: %s, want: %s.", idx, result, table.expected)
		}
	}
}
