package deadlock

import "sync"

/*
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	// ❌ Forgot to call wg.Done()
	wg.Wait() // 🔒 Blocks forever
}
*/
var wgd sync.WaitGroup

func Wg() {
	wgd.Add(1)
	go func() {
		defer wgd.Done()
		// work here
	}()
	wgd.Wait()

}
