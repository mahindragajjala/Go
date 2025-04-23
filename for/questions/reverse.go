package questions

import "fmt"

/*
Write a program to reverse a given number using a loop.
Let the number be 1234.
➡️ Expected Output: 4321

➡️ Goal: Learn how to rebuild a number digit-by-digit using %, /, and loop logic.
*/

func Reverse_Number_using_map() {
	GivenNumber := 1234
	Mapdata := []int{}
	oringial := 0
	for GivenNumber > 0 {
		Apart := GivenNumber % 10
		Mapdata = append(Mapdata, Apart)
		fmt.Println("a part..", Apart)
		GivenNumber = GivenNumber / 10
	}
	for _, value := range Mapdata {
		oringial = oringial*10 + value
	}
	fmt.Println(oringial)
}
func Without_map() {
	number := 1234
	original := number
	reversed := 0

	for number > 0 {
		digit := number % 10           // Extract the last digit
		reversed = reversed*10 + digit // Append digit to reversed number
		number /= 10                   // Remove the last digit
	}

	fmt.Printf("Original number: %d\n", original)
	fmt.Printf("Reversed number: %d\n", reversed)
}
