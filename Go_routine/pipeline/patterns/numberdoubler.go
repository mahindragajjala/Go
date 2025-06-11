package patterns

import (
	"fmt"
	"sync"
	"time"
)

/*
Number Doubler
   ➤ Read a list of integers and double each value using a pipeline.
   ➤ Stages: Generator → Doubler → Printer
*/

func Number_Doubler() {
	var wg sync.WaitGroup
	NormalNumber := make(chan int)
	DoublerNumber := make(chan int)
	Generator(&wg, NormalNumber)
	Doubler(&wg, NormalNumber, DoublerNumber)
	Printer(&wg, DoublerNumber)
	wg.Wait()
}
func Generator(wg *sync.WaitGroup, NormalNumber chan int) {
	wg.Add(1)
	go func() {
		for i := 0; i < 5; i++ {
			NormalNumber <- i
			time.Sleep(time.Second)
		}
		close(NormalNumber)
		defer wg.Done()
	}()

}
func Doubler(wg *sync.WaitGroup, number chan int, Double chan int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for data := range number {
			data = data * data
			Double <- data
			time.Sleep(time.Second)
		}
		close(Double)
	}()

}
func Printer(wg *sync.WaitGroup, double chan int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for data := range double {
			fmt.Println(data)
		}
	}()

}
