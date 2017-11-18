package simple

import "testing"

// TestSumNumbers basic table_driven test
// those "tables" are known in other languages as data providers
func TestSumNumbers(t *testing.T) {

	// we describe all test cases in a slice of objects
	// every object will contain the numbers we want to sum and the expected result
	table := []struct{
		a int
		b int
		expected int
	} {
		{
			a: 1,
			b: 1,
			expected: 2,
		},
		{
			a: 1,
			b: 9,
			expected: 10,
		},
		{
			a: 1,
			b: 2,
			expected: 3,
		},
	}

	// then iterate over the test cases and make assertion for each one
	for _, testCase := range table{
		result := SumNumbers(testCase.a, testCase.b)
		if result != testCase.expected {
			t.Errorf("Unexpected result! %d", result)
		}
	}
}

