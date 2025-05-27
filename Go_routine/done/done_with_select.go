package done

import (
	"fmt"
	"time"
)

func Workers() {
	var done = make(chan struct{})
	for i := 0; i < 5; i++ {
		go Worker(done, i)
	}
	time.Sleep(2 * time.Second)
	close(done)
}
func Worker(done chan struct{}, i int) {
	/* 	for {
		fmt.Printf("worker %v working...\n", i)
		time.Sleep(2 * time.Millisecond)
		break
	} */
	select {
	case <-done:
		fmt.Printf("worker %v is going to home\n", i)
		return
	default:
		fmt.Printf("worker %v working...\n", i)
		time.Sleep(200 * time.Millisecond)
	}
}
