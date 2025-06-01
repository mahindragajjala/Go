package problems

import "fmt"

/*
Design a struct Order with an embedded Customer struct and use method promotion to
print customer details from an Order instance.
*/
type CustomerData struct {
	Name string
	a    Addressdata
}
type Addressdata struct {
	City    string
	Pincode int
}

func Restaurant() {
	P1 := CustomerData{Name: "cherry", a: Addressdata{City: "hyderabad", Pincode: 500017}}
	P1.Method_promotion()
}

func (a *CustomerData) Method_promotion() {
	fmt.Println(a.Name)
	fmt.Println(a.a.City)
	fmt.Println(a.a.Pincode)
}
