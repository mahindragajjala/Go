package practice

import (
	"fmt"
)

// interface
type Interface_Data interface {
	Payment()
}
type Paypal struct {
	Name string
	id   int
	Ref  int
}

func (a Paypal) Payment() {
	fmt.Println("Payment on the :", a.Name)
}

type Stripe struct {
	id  int
	Ref int
}

func (b Stripe) Payment() {
	fmt.Println("Payment on the Stripe")
}

type Razorpay struct {
	id  int
	Ref int
}

func (c Razorpay) Payment() {
	fmt.Println("Payment on the Razorpay")
}

func Accessing_Interface(a Interface_Data) {
	a.Payment()
}

func Main_functionality() {
	paypal := Paypal{Name: "paypal", id: 123456789, Ref: 22222222}
	Accessing_Interface(paypal)
}
