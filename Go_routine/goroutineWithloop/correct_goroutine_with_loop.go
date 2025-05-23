package goroutinewithloop

import "fmt"

func Correct_goroutine_with_loop() {
	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
}
