package dtypes

/*
Polymorphism means "many forms".

Imagine you run an online store and offer multiple payment methods:
Credit Card
UPI
Net Banking
Cash on Delivery
All of these are different implementations of one action: Pay().

Imagine you have different types of vehicles:
Car 🚗
Bike 🏍️
Truck 🚚
Each vehicle can StartEngine(), but the internal mechanics differ.

Imagine you are creating a graphics app.
You deal with:
Circle 🔵
Rectangle ▭
Triangle 🔺
All have the same need: Calculate Area.

You want to save files, but the storage backend could be:
Local Disk 💽
Cloud (AWS S3, Google Cloud) ☁️
FTP Server 🌐

You want to send notifications to users:
Email 📧
SMS 📲
Push Notification 🔔

Scenario	                     Interface Concept	                                 Real-World Objects
Payment System	                 PaymentMethod.Pay()	                             UPI, Credit Card, NetBanking
Vehicles	                     Vehicle.StartEngine()	                             Car, Bike, Truck
Shapes	                         Shape.Area()	                                     Circle, Rectangle, Triangle
Employees	                     Employee.CalculateSalary()	                         FullTime, Freelancer, Intern
File Storage	                 Storage.Save()	                                     LocalDisk, Cloud, FTP
Notifications	                 Notifier.SendNotification()	                     Email, SMS, PushNotification

*/

import (
	"fmt"
)

// Define an interface
type PaymentMethod interface {
	Pay(amount float64)
}

// CreditCard struct
type CreditCard struct {
	CardNumber string
}

// Implement Pay() for CreditCard
func (cc CreditCard) Pay(amount float64) {
	fmt.Printf("Paid ₹%.2f using Credit Card ending with %s\n", amount, cc.CardNumber[len(cc.CardNumber)-4:])
}

// UPI struct
type UPI struct {
	UpiID string
}

// Implement Pay() for UPI
func (u UPI) Pay(amount float64) {
	fmt.Printf("Paid ₹%.2f via UPI ID: %s\n", amount, u.UpiID)
}

// NetBanking struct
type NetBanking struct {
	BankName string
}

// Implement Pay() for NetBanking
func (nb NetBanking) Pay(amount float64) {
	fmt.Printf("Paid ₹%.2f through NetBanking from %s Bank\n", amount, nb.BankName)
}

// Function that takes any PaymentMethod
func ProcessPayment(p PaymentMethod, amount float64) {
	p.Pay(amount) // Polymorphic call
}

func Polymorphism_function() {
	// Create payment method instances
	card := CreditCard{CardNumber: "1234-5678-9012-3456"}
	upi := UPI{UpiID: "user@upi"}
	netBank := NetBanking{BankName: "SBI"}

	// Process payments polymorphically
	ProcessPayment(card, 1500.00)
	ProcessPayment(upi, 500.50)
	ProcessPayment(netBank, 9999.99)

	/*



		user      paymenet() - card/upi/netbanl





	*/
}
