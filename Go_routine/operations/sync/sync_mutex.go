package sync

/*
SYNTAX:
var mu sync.Mutex

mu.Lock()
// critical section: only one goroutine can enter here
mu.Unlock()
*/

/* var mu sync.Mutex
var counter int

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mu.Lock()
		counter++
		mu.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go increment(&wg)
	go increment(&wg)

	wg.Wait()
	fmt.Println("Final Counter:", counter)
} */

/*
Always pair Lock() with Unlock() (use defer to avoid mistakes).
Don't call Unlock() if Lock() was not acquired (causes panic).
Ideal for simple mutual exclusion (1 reader/writer at a time).
*/
/*
------------------------------------------------------------------------------------------
var mu sync.Mutex
var counter int

Declares a global shared variable counter, which will be incremented by two concurrent goroutines.
This is the critical shared resource that must be protected using the mutex.
------------------------------------------------------------------------------------------
func increment(wg *sync.WaitGroup) {
	defer wg.Done()
This function will be run concurrently (via go increment(...)).
wg *sync.WaitGroup: Used to notify the main goroutine when this worker is done.
defer wg.Done(): Guarantees wg.Done() will be called when the function returns, decreasing the WaitGroup counter.
------------------------------------------------------------------------------------------
	for i := 0; i < 1000; i++ {
		mu.Lock()
		counter++
		mu.Unlock()
	}

	Loop: Repeats 1000 times.

mu.Lock(): Locks the mutex before accessing the counter.
counter++: Increments the global variable counter (critical section).
mu.Unlock(): Unlocks the mutex after updating.

✅ This ensures atomic updates to counter, avoiding race conditions.
------------------------------------------------------------------------------------------
func main() {
	var wg sync.WaitGroup
Entry point of the program.
Declares a WaitGroup wg to coordinate the two goroutines.
------------------------------------------------------------------------------------------
	wg.Add(2)
Adds 2 to the WaitGroup, indicating that we’re waiting for two tasks (goroutines) to finish.
------------------------------------------------------------------------------------------
	go increment(&wg)
	go increment(&wg)

Starts two goroutines running increment(&wg) concurrently.
Both share the same counter and mu.
------------------------------------------------------------------------------------------
	wg.Wait()
Main goroutine blocks here until the WaitGroup counter becomes 0.
This ensures we wait for both goroutines to finish before moving on.
------------------------------------------------------------------------------------------
*/
