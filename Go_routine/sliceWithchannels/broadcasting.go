package slicewithchannels

import (
	"fmt"
	"sync"
	"time"
)

func GoReceiver(i int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range ch {
		fmt.Printf("Receiver %d received: %d\n", i, val)
	}
}

func GoSender(sender chan int, receivers []chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range sender {
		for _, ch := range receivers {
			// send concurrently to avoid blocking
			go func(c chan int, v int) {
				c <- v
			}(ch, val)
		}
	}
}

func Broadcasting() {
	var wg sync.WaitGroup

	sender := make(chan int)

	receiverCount := 5
	receivers := make([]chan int, receiverCount)

	// Create and start receiver goroutines
	for i := 0; i < receiverCount; i++ {
		receivers[i] = make(chan int)
		wg.Add(1)
		go GoReceiver(i, receivers[i], &wg)
	}

	// Start sender goroutine
	wg.Add(1)
	go GoSender(sender, receivers, &wg)

	// Send some data to the sender channel
	for i := 1; i <= 5; i++ {
		fmt.Printf("Broadcasting value: %d\n", i)
		sender <- i
		time.Sleep(500 * time.Millisecond)
	}
	go func() {
		// Close sender channel
		close(sender)

		// Wait for broadcasting to finish before closing receiver channels
		wg.Wait()
	}()

	// Close all receiver channels
	for _, ch := range receivers {
		close(ch)
	}
}

/*
for val := range sender:
This loop receives values from the sender channel one by one until the channel is closed. The range over a channel blocks waiting for new values, so this will iterate over all values sent to sender.
*/
/*
for _, ch := range receivers:
This nested loop iterates over a slice (or array) of channels called receivers. Each ch is a channel that can receive int values.
*/
/*
go func(c chan int) { c <- val }(ch):
For each receiver channel ch, this spawns a new goroutine (a lightweight concurrent thread) which sends the current value val into the channel c (which is a parameter to the anonymous function).
The (ch) at the end passes the current ch as the argument c to the anonymous function. This avoids the common mistake of closure capturing the loop variable.
*/
/*
                           +-------------------+
                           |    main()         |
                           |-------------------|
                           | creates sender ch |
                           | creates receivers |
                           | starts goroutines |
                           +---------+---------+
                                     |
                                     v
                           +-------------------+
                           |   GoSender Goroutine
                           |-------------------|
                           |  for val in sender: |
                           |     broadcast val   |
                           +---------+-----------+
                                     |
         +---------------------------+---------------------------+
         |                           |                           |
         v                           v                           v
 +----------------+        +----------------+         +----------------+
 | Receiver 0 Goroutine |  | Receiver 1 Goroutine |   | Receiver 2 Goroutine |
 |----------------|      |----------------|       |----------------|
 |  <- ch[0]      |      |  <- ch[1]      |       |  <- ch[2]      |
 |  process val   |      |  process val   |       |  process val   |
 +----------------+      +----------------+       +----------------+

         ...
                        [ Similarly for Receiver 3, Receiver 4 ... ]


*/
