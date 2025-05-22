package multiplegoroutinesscheduling

import (
	"fmt"
	"sync"
)

/*
Write a program that prints even and odd numbers using two separate goroutines.
*/

// RANDOM PRINTING...
var wg sync.WaitGroup

func Even_Odd() {
	wg.Add(50)
	for i := 0; i < 50; i++ {
		if i%2 == 0 {
			go Even(i)
		} else {
			go Odd(i)
		}
	}
	wg.Wait()
}
func Even(i int) {
	defer wg.Done()
	fmt.Println("Even Number : ", i)
}
func Odd(i int) {
	defer wg.Done()
	fmt.Println("Odd Number : ", i)
}

// LINE BY LINE PRINTING FRIST ODD AFTER THAT EVEN
var wgdata sync.WaitGroup

func Printing_Line_By_Line() {
	num := 50

	wgdata.Add(2)
	go even(num)
	go odd(num)
	wgdata.Wait()

}
func even(n int) {
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			fmt.Println("Even", i)
		}
	}
	defer wgdata.Done()
}
func odd(n int) {
	for i := 1; i <= n; i++ {
		if i%2 != 0 {
			fmt.Println("Odd", i)
		}
	}
	defer wgdata.Done()
}

//SYNCRONIZED go routines..

func printEven(wg *sync.WaitGroup, max int) {
	defer wg.Done()
	for i := 0; i <= max; i++ {
		if i%2 == 0 {
			fmt.Printf("Even: %d\n", i)
		}
	}
}

func printOdd(wg *sync.WaitGroup, max int) {
	defer wg.Done()
	for i := 0; i <= max; i++ {
		if i%2 != 0 {
			fmt.Printf("Odd: %d\n", i)
		}
	}
}

func SYNCRONIZED() {
	var wg sync.WaitGroup
	max := 20

	wg.Add(2) // Exactly 2 goroutines

	go printEven(&wg, max)
	go printOdd(&wg, max)

	wg.Wait()
	fmt.Println("Finished printing.")
}
