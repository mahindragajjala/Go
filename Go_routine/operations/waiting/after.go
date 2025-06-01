package waiting

import "time"

func AFTER_GOROUTINE() {
	go func() {

		<-time.After(3 * time.Second) // blocks until timer fires

	}()

}
