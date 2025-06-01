package questions

import "fmt"

func Phonebook(name string) {
	directory := make(map[string]string)

	// Adding entries to the directory
	directory["Alice"] = "9876543210"
	directory["Bob"] = "9123456780"
	directory["Charlie"] = "9988776655"

	if no, MobileNo := directory[name]; MobileNo {
		fmt.Printf("USER IS FOUND : %v MOBILE NO : %v ", name, no)
	} else {
		fmt.Println("NOT FOUND")
	}
}
func Find_user() {
	var username string = "Bob"
	Phonebook(username)
}
