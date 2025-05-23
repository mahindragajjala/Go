package differenttypes

/*
Decoupling means reducing the dependency between different parts of your program.

When two parts are tightly coupled:
If you change one, you must change the other.
Difficult to maintain, extend, and test.

When decoupled:
Changes in one part don’t affect others.
Easier to maintain and extend.
*/

// Interface decouples the caller from the concrete implementation.

/* type PaymentMethod interface {
	Pay(amount float64)
}

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

// MakePayment doesn’t care about the payment type.
func MakePayment(p PaymentMethod, amount float64) {
	p.Pay(amount)
}

func main() {
	MakePayment(CreditCard{}, 500)
	MakePayment(DebitCard{}, 300)
	MakePayment(UPI{}, 150)
}
*/
