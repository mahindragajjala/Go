package waiting

import "time"

func SLEEP_GOROUTINE() {
	go func() {

		time.Sleep(2 * time.Second) // goroutine is sleeping

	}()

}
