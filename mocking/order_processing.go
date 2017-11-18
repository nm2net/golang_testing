package mocking

// Order object
type Order struct {
	id int
	isProcessed bool
	userEmail string
}

// Notifier will be the service we use to notify the user
type Notifier interface {
	NotifyUser(email string) bool
}

// OrderNotifier implementation
type OrderNotifier struct {}
func (on *OrderNotifier) NotifyUser(email string) bool {
	// ..
	return true
}

// ProcessOrder function that we want to test
// notice we accept here the interface Notifier
// but later in the test we pass reference of the object (using &)
// this is a general rule - when a method accepts interface
// we have to pass a reference to the object we want to use
func ProcessOrder(o *Order, n Notifier) *Order {
	o.isProcessed = n.NotifyUser(o.userEmail)

	return o
}




