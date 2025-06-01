package practice

import (
	"fmt"
	"sync"
)

func OneByOne() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go Printing(i, &wg)
		wg.Wait()
	}

}
func Printing(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("go routine - ", i)
}
