package coupled

import "fmt"

//define interfaces
type PaymentProcessor interface {
	ProcessPayment(amount float64) bool
}

type Notifier interface {
	SendNotification(msg string)
}

//Implement Services
//Implement Concrete Services
type StripeService struct{}

func (s *StripeService) ProcessPayment(amount float64) bool {
	fmt.Println("Stripe: Processing payment of", amount)
	return true
}

type EmailService struct{}

func (e *EmailService) SendNotification(msg string) {
	fmt.Println("Email Sent:", msg)
}

//Inject Dependencies
type OrderHandler struct {
	paymentService      PaymentProcessor
	notificationService Notifier
}

func NewOrderHandler(payment PaymentProcessor, notifier Notifier) *OrderHandler {
	return &OrderHandler{
		paymentService:      payment,
		notificationService: notifier,
	}
}

func (o *OrderHandler) PlaceOrder(amount float64) {
	if o.paymentService.ProcessPayment(amount) {
		o.notificationService.SendNotification("Order placed successfully!")
	}
}

//usage
func Decoupling() {
	payment := &StripeService{}
	notifier := &EmailService{}

	orderHandler := NewOrderHandler(payment, notifier)
	orderHandler.PlaceOrder(100.0)
}
