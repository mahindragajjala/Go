package parentchild

import (
	"fmt"
	"time"
)

/*
if above parent goroutine stops the child
*/
func ParentGoroutine() {
	go func() { //parent go routine
		for {
			select {
			default:
				fmt.Println("doing work!")
			}
		}
	}()
	time.Sleep(time.Second * 10)
	// time.Sleep(time.Hour * 299) // for long leave application
}
