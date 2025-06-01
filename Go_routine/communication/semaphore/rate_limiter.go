package semaphore

import (
	"fmt"
	"time"
)

/*
Rate Limiter Using Semaphore

Task:
Allow only 5 requests per second using a semaphore. Each request is handled by a goroutine. If 5 tokens are used, new requests wait until tokens are freed.

Goal:
Throttle request handling using a counting semaphore.

*/
/*
You want to throttle incoming requests so that only 5 requests are processed per second. If more than 5 requests come in at the same time, they should wait until the current requests finish and the "slots" (tokens) become available again.
*/
func Throttling() {
	channel := make(chan struct{}, 5)

	for i := 0; i <= 10; i++ {
		go Rate_Limiter(i, channel)
	}
	time.Sleep(3 * time.Second)
}
func Rate_Limiter(i int, channel chan struct{}) {
	channel <- struct{}{}
	fmt.Println("rate limited", i)
	time.Sleep(1 * time.Second)
	<-channel
}
