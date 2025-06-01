package headrecursion

import "fmt"

func DecimalToBinary(n int) {
	if n == 0 {
		return
	}
	DecimalToBinary(n / 2)
	fmt.Print(n % 2)
}
