package onerouterwithmultiplerouters

import (
	"fmt"
	"sync"
	"time"
)

func Sender_Values() {
	sender := make(chan []int)
	var wg sync.WaitGroup

	channels := make([]chan string, 2)
	for i := 0; i < len(channels); i++ {
		channels[i] = make(chan string)
		go func(ch chan string, id int) {
			for msg := range ch {
				fmt.Printf("Channel %d received: %s\n", id, msg)
			}
		}(channels[i], i)
	}

	// Sender goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		sender <- []int{1, 2, 3}
		time.Sleep(time.Second)
		close(sender)
	}()

	// Processor goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for data := range sender {
			for i := 0; i < len(channels); i++ {
				wg.Add(2)
				go Addition(&wg, data, channels[i])
				go Multiplication(&wg, data, channels[i])
			}
		}
	}()

	wg.Wait()

	// Close all channels
	for _, ch := range channels {
		close(ch)
	}
}

func Addition(wg *sync.WaitGroup, data []int, ch chan string) {
	defer wg.Done()
	sum := 0
	for _, v := range data {
		sum += v
	}
	ch <- fmt.Sprintf("Addition result: %d", sum)
}

func Multiplication(wg *sync.WaitGroup, data []int, ch chan string) {
	defer wg.Done()
	product := 1
	for _, v := range data {
		product *= v
	}
	ch <- fmt.Sprintf("Multiplication result: %d", product)
}
