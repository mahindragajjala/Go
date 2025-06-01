package deadlock

/*
var mu1, mu2 sync.Mutex

func main() {
	go func() {
		mu1.Lock()
		time.Sleep(1 * time.Second)
		mu2.Lock() // 🔒 Waits for mu2
	}()

	go func() {
		mu2.Lock()
		time.Sleep(1 * time.Second)
		mu1.Lock() // 🔒 Waits for mu1
	}()

	time.Sleep(3 * time.Second)
}

*/
/* // In all goroutines:
mu1.Lock()
mu2.Lock()
// ...
mu2.Unlock()
mu1.Unlock() */
