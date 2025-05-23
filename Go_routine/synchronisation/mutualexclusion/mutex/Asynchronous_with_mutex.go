package mutex

import (
	"fmt"
	"sync"
	"time"
)

/*
Asynchronous means:

Start a task and keep going without waiting for it to finish."
*/
/*
In asynchronous systems, different parts of the program may work independently, but mutex can still be used to synchronize access to shared data.

You could call it asynchronous execution with synchronous locking.
*/

func Asynchronous() {
	fmt.Println("ASYNCHRONOUS--->")
	var mu sync.Mutex
	data := make(map[string]int)

	var wg sync.WaitGroup
	keys := []string{"a", "b", "c"}

	for _, key := range keys {
		fmt.Printf("go routine %v started..\n", key)
		wg.Add(1)
		go func(k string) {
			defer wg.Done()
			time.Sleep(1 * time.Second) // simulate async work // 🟣 Simulates async/delayed work
			mu.Lock()                   // 🟡 SYNC: Protects access to map
			data[k]++
			mu.Unlock()
		}(key)
		fmt.Printf("go routine %v ended..\n", key)
	}

	wg.Wait()
	mu.Lock()
	fmt.Println("Data map:", data)
	mu.Unlock()
}

/*
| Aspect          | Synchronous                       | Asynchronous                           |
| --------------- | --------------------------------- | -------------------------------------- |
| Execution Style | Linear (one-by-one)               | Concurrent (multiple goroutines)       |
| Mutex Usage     | Not always needed                 | Needed to protect shared data          |
| Mutex Behavior  | Blocks next operation till unlock | Makes async operations thread-safe     |
| Risk            | Rare race condition               | High race condition risk without mutex |

*/

/*
func main() {
    var mu sync.Mutex
    data := make(map[string]int)

mu is the mutex that will ensure synchronized access to data.
data is the shared map that goroutines will write to.
At this point, the program is still running synchronously — step-by-step in main().
*/

/*
    var wg sync.WaitGroup
    keys := []string{"a", "b", "c"}

wg is used to wait for all goroutines to finish.
keys is a slice of strings — these are used to spawn 3 goroutines.
Still synchronous.
*/
/*
    for _, key := range keys {
        wg.Add(1)
        go func(k string) {
            defer wg.Done()
            time.Sleep(1 * time.Second) // Simulate async work
            mu.Lock()
            data[k]++
            mu.Unlock()
        }(key)
    }

	for _, key := range keys {
Loop over each key: "a", "b", "c".

wg.Add(1)
Tell the WaitGroup that we’re starting a goroutine.

go func(k string) { ... }(key)
🚀 Asynchronous Launch Point:

A new goroutine is started here.
func(k string) is run concurrently (in the background).
Execution of main() does NOT wait for this function to finish.
Instead, all 3 goroutines are launched almost instantly in parallel.
*/
/*
Inside the goroutine:

defer wg.Done()
This line schedules wg.Done() to be called when the goroutine finishes.
It signals that the goroutine's work is complete.
*/
/*
time.Sleep(1 * time.Second)
"Simulated Delay / Asynchronous Wait":
Goroutine sleeps independently.
During this time, other goroutines or the main thread can run freely.
This is not blocking the main() thread or other goroutines.
This is the essence of asynchronous behavior — things happen independently.

*/
/*
mu.Lock()
 Synchronous Entry to Critical Section:

After sleep, the goroutine waits for the mutex to be unlocked.
If another goroutine is already inside the mutex, it will wait (block) here.
This part becomes synchronous (one at a time) despite the async start.
*/
/*
data[k]++

Safely increments the value for the key (e.g., "a", "b", or "c").
Only one goroutine can do this at a time, due to the mutex.
*/
/*
Time (approx) ➝   0s        1s             1.1s           1.2s           1.3s
                ─────────────────────────────────────────────────────────────────

Main Thread     | Launch G1, G2, G3   | Waits at wg.Wait()

Goroutine G1    | Sleep ─────────────▶| mu.Lock() ✅
                                     | data["a"]++
                                     | mu.Unlock()

Goroutine G2    | Sleep ─────────────▶| mu.Lock() ⏳ (waits for G1)
                                                  | mu.Lock() ✅
                                                  | data["b"]++
                                                  | mu.Unlock()

Goroutine G3    | Sleep ─────────────▶| mu.Lock() ⏳ (waits for G2)
                                                               | mu.Lock() ✅
                                                               | data["c"]++
                                                               | mu.Unlock()

Main Thread     |                               | Waits         | Waits       | Proceeds ➝ Prints data

*/
