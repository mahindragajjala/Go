package stringtypes

import "fmt"

func Strings() {
	//Declare
	var Stringdata string = "Hey this is string data"
	fmt.Println(Stringdata)

	//Accessing
	fmt.Println(Stringdata[0])

	//Immutable
	//Stringdata[0] = P //error will get

	//Find address of String
	fmt.Println(&Stringdata)

	//A string is a slice of bytes

	str := "Hello André!"
	for i := range 13 {
		fmt.Print(str[i], " ")
	}
	fmt.Println()
	//72 101 108 108 111 32 65 110 100 114 195 169 33
	//These numbers represent bytes in decimal notation.
	// Let’s transform these bytes in hexadecimal by using
	// the %x verb:
	for i := range 13 {
		fmt.Printf("%x ", str[i])
	}
	fmt.Println()

	//type checking
	var stringdata string = "TYPE"
	fmt.Printf("TYPE CHECKING : %T \n ", stringdata)
	fmt.Printf("A PART : %T", stringdata[0])

	// slice of bytes
	var SliceOfBytes string = "TYPE"
	b := []byte(SliceOfBytes)
	fmt.Println("Length of byte slice:", len(b))
}
