package multiplegoroutinesscheduling

import (
	"fmt"
	"sync"
)

/*
Write a function that spawns 10 goroutines, each computing the square of a number.
*/

func Spawning() {
	var wgs sync.WaitGroup
	wgs.Add(10)
	for i := 1; i <= 10; i++ {
		go spawn(i, &wgs)
	}
	wgs.Wait()
}
func spawn(i int, wgc *sync.WaitGroup) {
	defer wgc.Done()
	fmt.Println(i * i)
}

func Sync_Swawning() {
	var wgs sync.WaitGroup
	for i := 0; i < 10; i++ {
		wgs.Add(1)
		go func(i int) {
			fmt.Println(i * i)
			defer wgs.Done()
		}(i)
		wgs.Wait()
	}

}
