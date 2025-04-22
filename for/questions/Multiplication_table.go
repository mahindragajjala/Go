package questions

import "fmt"

func Multiple_questions(a int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%v * %v = %v \n", a, i, a*i)
	}
}
func Multiple_questions_userdata() {
	Multiple_questions(2)
}
