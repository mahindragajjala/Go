package questions

import "math"

var Prime_number []int

func PrimeNumber_checking() {
	data := []int{2, 3, 4, 5, 6, 7, 8, 9}
	for _, value := range data {
		isPrime(value)
	}

}
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrtN; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
