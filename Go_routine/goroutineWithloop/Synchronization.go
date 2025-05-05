package goroutinewithloop

import (
	"fmt"
	"sync"
)

func Synchronization() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(i)
		}()
	}

	wg.Wait()

}
