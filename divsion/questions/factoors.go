package questions

import "fmt"

func findFactors(n int) []int {
	var factors []int
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}

func Findfactors_input() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	factors := findFactors(n)
	fmt.Printf("The factors of %d are: ", n)
	for _, factor := range factors {
		fmt.Printf("%d ", factor)
	}
	fmt.Println()
}
