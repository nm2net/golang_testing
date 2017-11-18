package sub_tests

import "testing"

// TestSumNumbers more elegant table_driven test including sub tests
// those "tables" are known in other languages as data providers
func TestSumNumbers(t *testing.T) {

	// we describe all test cases in a map of string:object
	// every element of the map will contain the title of th test
	// and the object containing the expected result and the numbers we want to sum
	table := map[string]struct{
		a int
		b int
		expected int
	} {
		"test 1": {
			a: 1,
			b: 1,
			expected: 2,
		},
		"test 2": {
			a: 1,
			b: 9,
			expected: 10,
		},
		"test C": {
			a: 1,
			b: 2,
			expected: 3,
		},
	}

	// then iterate over the test cases and run sub test for each one
	for testName, testCase := range table{
		t.Run(testName, func(t *testing.T) {
			result := SumNumbers(testCase.a, testCase.b)

			if result != testCase.expected {
				t.Errorf("Unexpected result! %d", result)
			}
		})
	}
}
