package coupled

import "fmt"

/*

Scenario: You are writing a simple OrderService for an online store.
When the user places an order, the service must check the product stock and
then confirm the order.

In this version, the OrderService should create its own InventoryChecker directly
inside the method.

✅ Task:

Write a Go program where OrderService directly creates and uses InventoryChecker
inside the PlaceOrder method.

The InventoryChecker will have a CheckStock method that always returns true.

*/
// type PlaceOrder struct {
// 	a InventoryChecker
// 	b OrderService
// }
// type InventoryChecker struct{}

// func (i *InventoryChecker) CheckStock(product string) bool {
// 	// hardcoded logic for now
// 	return true
// }

// type OrderService struct{}

// func (o *OrderService) PlaceOrder(product string) {
// 	fmt.Println("")
// }
// func (a *PlaceOrder) Placeorder(data string) {
// 	if a.a.CheckStock(data) {
// 		a.b.PlaceOrder(data)
// 	}
// }
// func PLACE_ORDER() {
// 	//var a = InventoryChecker{}
// 	var b = OrderService{}
// 	b.PlaceOrder("data")
// }

/*
Scenario: Now rewrite the same example using interfaces and dependency injection to
decouple the OrderService from the InventoryChecker.

This time, OrderService should receive an interface rather than directly creating an
InventoryChecker.

✅ Task:

Create an InventoryService interface with a method CheckStock(product string) bool.

Implement this interface in a concrete type MyInventoryChecker.

Inject this service into OrderService via its constructor.

In main(), create the MyInventoryChecker and inject it into OrderService.
*/
var ProductsCount int

type Order_Service interface {
}

type ProductAvailability struct {
	Products int
}

func (c *ProductAvailability) CheckProducts() bool {
	if c.Products == 0 {
		fmt.Println("Products not available...")
		return false
	} else {
		return true
	}
}

type Order struct {
	Yes bool
	No  bool
}

func (c *Order) ConfirmSellProduct() {
	fmt.Println("Order confirmed!")
}
