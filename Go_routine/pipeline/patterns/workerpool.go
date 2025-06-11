package patterns

import (
	"fmt"
	"sync"
	"time"
)

func WorkerPool() {
	var wg sync.WaitGroup
	generatorCh := make(chan int)
	processCh := make(chan int)

	WorkerPool_Generator(generatorCh, &wg)
	WorkerPool_Process(generatorCh, processCh, &wg)
	WorkerPool_Printer(processCh, &wg)

	wg.Wait()
}

func WorkerPool_Generator(ch chan int, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 25; i++ {
			ch <- i
		}
		close(ch)
	}()
}

func worker(input int, sem chan struct{}, output chan int, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		sem <- struct{}{}       // acquire token
		time.Sleep(time.Second) // simulate heavy work
		output <- input * input // process
		<-sem                   // release token
	}()
}

func WorkerPool_Process(inputCh chan int, outputCh chan int, wg *sync.WaitGroup) {
	go func() {
		sem := make(chan struct{}, 3) // limit to 3 concurrent workers

		for v := range inputCh {
			worker(v, sem, outputCh, wg)
		}

		// Wait for all workers to finish before closing output channel
		go func() {
			wg.Wait()
			close(outputCh)
		}()
	}()
}

func WorkerPool_Printer(outputCh chan int, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range outputCh {
			fmt.Println(v)
		}
	}()
}
