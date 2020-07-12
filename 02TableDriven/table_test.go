package table

import "testing"

func TestTableCalculate(t *testing.T) {
	var testData = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{3, 9},
		{0, 0},
		{-2, 4},
		{112, 12544},
	}

	for _, test := range testData {
		if output := tablePow(test.input); output != test.expected {
			t.Errorf("Test Failed: %d inputted, %d expected, recieved: %d", test.input, test.expected, output)
		} else {
			t.Log("Passed")
		}
	}
}
