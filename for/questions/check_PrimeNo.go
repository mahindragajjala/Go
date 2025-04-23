package questions

import "fmt"

/* func Checking_Prime_No(p int) bool {
	if p > 1 {
		a := p % p
		b := p % 1
		if a == 0 && b == 0 {
			return true
		}
	} else {
		fmt.Println("given number is not a prime")
		return false
	}
	return false
} */
func Checking_Prime_No(p int) bool {
	if p <= 1 {
		fmt.Println("given number is not a prime")
		return false
	}
	for i := 2; i*i <= p; i++ {
		if p%i == 0 {
			return false
		}
	}
	return true
}

/*
Input: An integer p

If p <= 1
→ Print "given number is not a prime"
→ Return false

Loop i from 2 to √p:

If p % i == 0
→ Return false (Not a prime)

If no divisors found, return true (It is a prime)
*/
func Checking_Prime_No_userinput() {
	if Checking_Prime_No(7) {
		fmt.Println("PRIME")
	} else {
		fmt.Println("NOT PRIME")
	}
}
