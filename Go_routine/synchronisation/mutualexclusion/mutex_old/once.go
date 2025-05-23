package mutex

import (
	"fmt"
	"sync"
)

var once sync.Once

func Synchronization_once() {
	for i := 0; i < 5; i++ {
		go func() {
			once.Do(Synchro)
		}()
	}
	fmt.Scanln()
}
func Synchro() {
	fmt.Println("checking sync.once - it will print only once ")
}
