package coupled

import "fmt"

/*
In a coupled system, structs are directly dependent on each other.
If one struct changes, the other often must change too.
This creates tight binding.
*/
type PaymentService struct{}

func (p *PaymentService) ProcessPayment(amount float64) bool {
	fmt.Println("Processing payment:", amount)
	return true
}

type NotificationService struct{}

func (n *NotificationService) SendNotification(msg string) {
	fmt.Println("Sending Notification:", msg)
}

type Coupled_OrderHandler struct {
	paymentService      PaymentService
	notificationService NotificationService
}

func (o *Coupled_OrderHandler) Coupled_PlaceOrder(amount float64) {
	if o.paymentService.ProcessPayment(amount) {
		o.notificationService.SendNotification("Order placed successfully!")
	}
}

func Coupled_Data() {
	d := Coupled_OrderHandler{}
	d.Coupled_PlaceOrder(5665)
}

/*
❌ Problems:
Coupled_OrderHandler is directly creating and depending on concrete implementations
(PaymentService, NotificationService).

Hard to test (you can’t mock services easily).

Hard to change (if Stripe is replaced by Razorpay — you'd edit everywhere).
*/
