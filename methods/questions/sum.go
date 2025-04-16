package questions

import "fmt"

var sum int

func Main_sum() {
	numbers := []int{2, 3, 4, 5, 6}
	for _, value := range numbers {
		Sum(&value)
	}
	fmt.Println(sum)
}
func Sum(a *int) {
	sum = sum + *a
}
