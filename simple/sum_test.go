package simple

import "testing"

// TestSumNumbers basic test
func TestSumNumbers(t *testing.T) {
	result := SumNumbers(1, 2)

	if result != 3 {
		t.Error("Unexpected result")
	}
}
