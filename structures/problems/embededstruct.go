package problems

import "fmt"

/*
Define two structs, Address and Customer, where Address is embedded in Customer.
Initialize Customer and print both Name and City.

*/
type Address struct {
	City string
}

type Customer struct {
	Name string
	d    Address
}

func Embeded_struct() {
	cus1 := Customer{Name: "sai", d: Address{City: "hyderabad"}}
	fmt.Println(cus1.Name)
	fmt.Println(cus1.d.City)
}
