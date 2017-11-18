package simple

import "testing"

// We have to overwrite the TestMain method
// It is being called when the test file is executed
// Accepts testing.M object from the testing framework,
// that we will use to execute al tests in this file with m.Run()
// that way having control of what happens before and after all tests are executed
func TestMain(m *testing.M) {
	println("do stuff before all tests")
	m.Run()
	println("do stuff after all tests")
}

// TestSumNumbers basic test
func TestSumNumbers(t *testing.T) {
	result := SumNumbers(1, 2)

	if result != 3 {
		t.Error("Unexpected result")
	}
}
