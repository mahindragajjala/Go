package pipeline

import (
	"fmt"
	"sync"
)

// Send values from one input channel to multiple output channels
func fanOut(input <-chan int, outputs []chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range input {
		for _, out := range outputs {
			out <- val
		}
	}
	for _, out := range outputs {
		close(out)
	}
}

func receive(ch <-chan int, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range ch {
		fmt.Printf("Receiver %d got: %d\n", id, val)
	}
}

func ChannelToChannels() {
	input := make(chan int)
	var wg sync.WaitGroup

	// Create multiple output channels
	numReceivers := 3
	outputs := make([]chan int, numReceivers)
	for i := range outputs {
		outputs[i] = make(chan int)
	}

	// Start fan-out
	wg.Add(1)
	go fanOut(input, outputs, &wg)

	// Start receivers
	for i, ch := range outputs {
		wg.Add(1)
		go receive(ch, i+1, &wg)
	}

	// Send data into input channel
	for i := 0; i < 5; i++ {
		input <- i
	}
	close(input)

	wg.Wait()
}

/*
┌──────────────────────────────┐
│         MAIN THREAD          │
└────────────┬─────────────────┘
             │
             ▼
     ┌─────────────────┐
     │ Create input ch │◄────────────┐
     └─────────────────┘             │
             │                       │
     ┌────────────────────┐         │
     │ Create output chans│◄────┐   │
     └────────────────────┘     │   │
             │                  │   │
     ┌────────────────────┐     │   │
     │ Start fanOut goroutine ──┼───┘
     └────────────────────┘     │
             │                  │
     ┌────────────────────┐     │
     │ Start receiver #1   │◄───┤
     └────────────────────┘     │
     ┌────────────────────┐     │
     │ Start receiver #2   │◄───┤
     └────────────────────┘     │
     ┌────────────────────┐     │
     │ Start receiver #3   │◄───┘
     └────────────────────┘
             │
             ▼
     ┌──────────────────────┐
     │ Send values to input │
     └──────────────────────┘
             │
             ▼
     ┌────────────────────────────┐
     │      fanOut(input)         │ ◄────────────────────────────┐
     │ (reads from input channel) │                              │
     │  for each value:           │                              │
     │   ┌────────────┐           │                              │
     │   │ out[0] ◄── val ◄────┐  │                              │
     │   └────────────┘        │  │                              │
     │   ┌────────────┐        │  │   Sends to all output chans │
     │   │ out[1] ◄── val ◄────┼──┼─────────────────────────────► Receivers
     │   └────────────┘        │  │
     │   ┌────────────┐        │  │
     │   │ out[2] ◄── val ◄────┘  │
     │   └────────────┘           │
     │                            │
     │ After input closed:        │
     │  Close all outputs         │
     └────────────────────────────┘
             │
             ▼
     ┌────────────────────────────┐
     │    Receivers read data     │
     │  from output channels      │
     │  e.g., Receiver 1:         │
     │     out[0] → print         │
     │  ... until channel closed  │
     └────────────────────────────┘
             │
             ▼
     ┌────────────────────────────┐
     │      wg.Wait()             │
     │ Wait for all to complete  │
     └────────────────────────────┘
             │
             ▼
     ┌────────────────────────────┐
     │     Program Ends           │
     └────────────────────────────┘


*/
