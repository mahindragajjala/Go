package commandpassing

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Function each goroutine will run
func worker(id int, stopChan chan struct{}) {
	fmt.Printf("Goroutine %d started\n", id)
	for {
		select {
		case <-stopChan:
			fmt.Printf("Goroutine %d stopping\n", id)
			return
		default:
			fmt.Printf("Goroutine %d is working...\n", id)
			time.Sleep(1 * time.Second)
		}
	}
}

func CommandPassing() {
	stopChans := make([]chan struct{}, 5)

	// Create and start 5 goroutines
	for i := 0; i < 5; i++ {
		stopChans[i] = make(chan struct{})
		go worker(i+1, stopChans[i])
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter goroutine number to stop (1-5), or 'exit' to quit: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("Exiting. Stopping all goroutines...")
			for i := 0; i < 5; i++ {
				close(stopChans[i])
			}
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil || num < 1 || num > 5 {
			fmt.Println("Invalid input. Please enter a number between 1 and 5 or 'exit'")
			continue
		}

		// Close the stop channel for the selected goroutine
		select {
		case <-stopChans[num-1]: // already closed
			fmt.Printf("Goroutine %d is already stopped.\n", num)
		default:
			close(stopChans[num-1])
		}
	}

	// Wait for a moment to let all goroutines print their stop messages
	time.Sleep(2 * time.Second)
	fmt.Println("Program finished.")
}
