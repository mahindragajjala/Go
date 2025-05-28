package patterns

import "fmt"

/*
Number Doubler
   ➤ Read a list of integers and double each value using a pipeline.
   ➤ Stages: Generator → Doubler → Printer
*/

func Number_Doubler() {
	number := Generator()
	double := Doubler(number)
	Printer(double)
}

func Generator() chan int {
	Number := make(chan int)
	for i := 0; i < 5; i++ {
		Number <- i
	}
	return Number
}

/*
Generator
❌ This is writing to the channel (Number <- i) without a concurrent reader, so this causes a deadlock.

✅ Fix: Run a goroutine to write values and close the channel when done.
*/
func Doubler(num chan int) chan int {
	ch := make(chan int)
	digit := <-num
	digit = digit * digit
	ch <- digit
	return ch
}
func Printer(double chan int) {
	fmt.Println(<-double)
}
