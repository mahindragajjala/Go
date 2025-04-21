package questions

import "fmt"

/*
Create a map of country names and their capitals.

Write a function to fetch the capital for a given country.
*/
/* var Country = make(map[string]string)

func Search_value(countrydata string) {
	Country["India"] = "+91"
	Country["United States"] = "+1"
	Country["United Kingdom"] = "+44"
	Country["Australia"] = "+61"
	Country["Canada"] = "+1"
	for key, value := range Country {
		if countrydata == key {
			fmt.Println(value)
		}
	}
}
func Search_value_country() {

	Search_value("+91")
	Search_value("+1")
}
*/
var Country = make(map[string]string)

func init() {
	Country["India"] = "New Delhi"
	Country["United States"] = "Washington, D.C."
	Country["United Kingdom"] = "London"
	Country["Australia"] = "Canberra"
	Country["Canada"] = "Ottawa"
}

func GetCapital(country string) {

	//This tries to fetch the value from the map using the country as a key.

	if capital, exists := Country[country]; exists {
		fmt.Printf("The capital of %s is %s\n", country, capital)
	} else {
		fmt.Println("Country not found!")
	}
}

func Search_value_country() {
	GetCapital("India")
	GetCapital("United States")
	GetCapital("France")
}
