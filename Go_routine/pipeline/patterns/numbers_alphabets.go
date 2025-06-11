package patterns

import (
	"fmt"
	"strconv"
	"sync"
)

func Numbers_albhabets() {
	//using Print with loops

	//nums
	for i := 1; i <= 9; i++ {
		fmt.Print(i)
	}
	fmt.Println()
	//alphabets
	alphabets := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}
	for i := 0; i < len(alphabets); i++ {
		fmt.Print(alphabets[i])
	}
	fmt.Println()
	Intergrate(alphabets)
}
func Intergrate(alphabets []string) {
	for i := 1; i <= 9; i++ {
		fmt.Println(i, alphabets[i-1])
	}
}

/*
output:
123456789
abcdefghi
	1a
	2b
	3c
	4d
	5e
	6f
	7g
	8h
	9i
*/

// Using concurrency

func Concurrency() {
	var wg sync.WaitGroup
	numbers := make(chan int)
	alphabets := make(chan string)

	wg.Add(3)
	go Numbers(numbers, &wg)
	go Alphabets(numbers, alphabets, &wg)
	go Printing(alphabets, &wg)
	wg.Wait()
}

func Numbers(numbers chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 9; i++ {
		fmt.Print(i)
		numbers <- i
	}
	close(numbers)
	fmt.Println()
}

func Alphabets(numbers chan int, alphabets chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	chars := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}
	for i := 0; i < 9; i++ {
		fmt.Print(chars[i])
	}
	fmt.Println()

	i := 0
	for num := range numbers {
		result := "\t" + strconv.Itoa(num) + chars[i]
		alphabets <- result
		i++
	}
	close(alphabets)
}

func Printing(alphabets chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range alphabets {
		fmt.Println(val)
	}
}
