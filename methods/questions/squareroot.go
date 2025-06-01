package questions

import (
	"fmt"
	"math"
)

func SquareRoot() {
	data := []int{2, 3, 4, 5, 6, 7, 8, 9}
	for _, value := range data {
		sqrtN := int(math.Sqrt(float64(value)))
		fmt.Println(sqrtN)
	}

}
