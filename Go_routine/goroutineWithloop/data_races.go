package goroutinewithloop

import "sync"

/*
DATA RACES
Problem: Goroutines accessing and modifying shared data concurrently.


*/
func Data_Races() {
	count := 0
	for i := 0; i < 5; i++ {
		go func() {
			count++
		}()
	}
	/*
	   This causes a race condition and undefined behavior.
	*/
}

/*
✅ Fix: Use mutex or atomic operations.
*/
func Mutex_for_raceconditions() {
	var mu sync.Mutex
	count := 0
	for i := 0; i < 5; i++ {
		go func() {
			mu.Lock()
			count++
			mu.Unlock()
		}()
	}
}
