package closechannel

import "fmt"

//Using a done or quit channel (Graceful Shutdown)

func Using_Done_quiet() {
	data := make(chan int)
	done := make(chan struct{})

	go func() {
		for i := 0; i < 5; i++ {
			data <- i
		}
		close(done) // signal completion
	}()

	for {
		select {
		case val := <-data:
			fmt.Println("Received:", val)
		case <-done:
			fmt.Println("All done.")
			return
		}
	}

}
