package checking

import "fmt"

func Divisor() {
	number := 50
	var Divisor []int
	for i := 1; i <= 50; i++ {
		if number%i == 0 {
			Divisor = append(Divisor, i)
		}
	}
	fmt.Println(Divisor)
}
