package directions

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
Build a generator stage that emits numbers.
Pass the output to a function that doubles the values and prints them.
*/
func generator(channel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var count int = 5
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("Generating %d random numbers:\n", count)
	for i := 0; i < count; i++ {
		// fmt.Println(rand.Intn(count))
		channel <- rand.Intn(count)
		fmt.Println("sent...")
		time.Sleep(time.Second)
	}
	close(channel)
}
func doubles(channel chan int, wg *sync.WaitGroup, multiplechannel chan int) {
	defer wg.Done()
	for data := range channel { //receiver
		fmt.Println("received", data)
		multiplechannel <- data * data //sender
		time.Sleep(time.Second)
	}
	close(multiplechannel)
}
func Squares(wg *sync.WaitGroup, multiplechannel chan int) {
	defer wg.Done()
	for data := range multiplechannel {
		fmt.Println(data)
		time.Sleep(time.Second)
	}
}
func Simple_pipeline() {
	channel := make(chan int)
	multiplechannel := make(chan int)
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		generator(channel, &wg)
	}()
	go func() {
		doubles(channel, &wg, multiplechannel)
	}()
	go func() {
		Squares(&wg, multiplechannel)
	}()
	wg.Wait()
}
