package mocking

import (
	"testing"
)

// OrderNotifierMock is is the object we will use to mock OrderNotifier by implementing Notifier interface
// It has FakeNotifyUserResponse property that we are going to use
// to fake it's response when notifyUser method is called
type OrderNotifierMock struct {
	FakeNotifyUserResponse bool
}
// NotifyUser has same signature as the interface
func (om *OrderNotifierMock) NotifyUser(email string) bool {
	return om.FakeNotifyUserResponse
}

func TestProcessOrder(t *testing.T) {
	table := map[string]struct{
		notifierResult bool
		expectedResult bool
	} {
		"failed response": {
			notifierResult: false,
			expectedResult: false,
		},
		"successful response": {
			notifierResult: true,
			expectedResult: true,
		},
	}

	for testName, testCase := range table {
		t.Run(testName, func(t *testing.T) {

			notifier := &OrderNotifierMock{
				// this is where we manipulate the response of notifyUser method
				FakeNotifyUserResponse: testCase.notifierResult,
			}
			order :=  &Order{
				id: 123,
				userEmail: "niko@email.com",
			}
			result := ProcessOrder(order, notifier)

			if result.isProcessed != testCase.expectedResult {
				t.Errorf("Unexpected result %v", result.isProcessed)
			}
		})
	}
}
