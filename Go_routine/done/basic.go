package done

import (
	"fmt"
	"time"
)

/*
In Go, the close keyword is used with channels to indicate that "no more values will be sent on that channel".

It is not used with files, network connections, or other resources (for those, you use methods like .Close() from types that implement io.Closer).
*/
func Childrens() {
	var done = make(chan struct{})
	for i := 0; i < 5; i++ {
		go Children(done, i)
	}
	time.Sleep(2 * time.Second)
	close(done)
}
func Children(done chan struct{}, i int) {
	for {
		fmt.Printf("children %v playing...\n", i)
		time.Sleep(2 * time.Millisecond)
		break
	}
	fmt.Printf("children %v is going to home\n", i)
}
