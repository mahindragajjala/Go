package differenttypes

import "fmt"

/*
Imagine you’re building an online shopping platform (like Amazon).
You allow customers to pay in different ways:

Debit Card
Credit Card
UPI
Wallet (Paytm, PhonePe, etc.)

All these payment methods have the same action: Pay(), but the logic is different
for each one

*/
// Step 1: Define an interface
type PaymentMethod interface {
	Pay(amount float64)
}

// Step 2: Implement different payment types
type CreditCard struct{}
type DebitCard struct{}
type UPI struct{}

func (c CreditCard) Pay(amount float64) {
	fmt.Printf("Paid ₹%.2f using Credit Card.\n", amount)
}

func (d DebitCard) Pay(amount float64) {
	fmt.Printf("Paid ₹%.2f using Debit Card.\n", amount)
}

func (u UPI) Pay(amount float64) {
	fmt.Printf("Paid ₹%.2f using UPI.\n", amount)
}

// Step 3: Function that uses polymorphism
func MakePayment(p PaymentMethod, amount float64) {
	p.Pay(amount)
}

func Polymorphism() {
	// Using different payment methods
	var payment PaymentMethod

	payment = CreditCard{}
	MakePayment(payment, 500.00)

	payment = DebitCard{}
	MakePayment(payment, 300.50)

	payment = UPI{}
	MakePayment(payment, 150.75)
}
