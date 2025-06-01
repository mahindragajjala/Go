package routing

import (
	"fmt"
	"time"
)

func Concurrency() {
	go goroutine1()
	go goroutine2()

	// Give goroutines time to finish
	time.Sleep(2 * time.Second)
}

func goroutine1() {
	for i := 1; i <= 5; i++ {
		fmt.Println("goroutine 1:", i)
		time.Sleep(200 * time.Millisecond)
	}
	fmt.Println("GOROUTINE 1 DONE")
}

func goroutine2() {
	for i := 1; i <= 5; i++ {
		fmt.Println("goroutine 2:", i)
		time.Sleep(200 * time.Millisecond)
	}
	fmt.Println("GOROUTINE 2 DONE")
}

func Concurrency_timing() {
	go Routine1()
	go Routine2()
	time.Sleep(5 * time.Second)
}
func Routine1() {
	for {
		now := time.Now()
		fmt.Println("GO ROUTINE 1")
		fmt.Println(time.Since(now))
		time.Sleep(600 * time.Millisecond)
	}

}
func Routine2() {
	for {
		now := time.Now()
		fmt.Println("GO ROUTINE 2")
		fmt.Println(time.Since(now))
		time.Sleep(600 * time.Millisecond)
	}

}
