package codingquestions

import (
	"fmt"
	"sync"
)

type Details struct {
	Data int
	mu   sync.Mutex
}

func Synchronisation_function() {
	var D Details
	D.SYNCHRONISATION_MUTEX()
}
func (D *Details) SYNCHRONISATION_MUTEX() {
	D.mu.Lock()
	for i := 0; i < 50; i++ {
		D.Data++
	}
	fmt.Println(D.Data)
	D.mu.Unlock()
}
