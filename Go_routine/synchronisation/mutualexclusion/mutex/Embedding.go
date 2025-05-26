package mutex

import (
	"fmt"
	"sync"
)

type Embedding struct {
	mu   sync.Mutex
	data int
}

func Goroutines() {
	var data = Embedding{
		data: 0,
	}
	var wg sync.WaitGroup
	wg.Add(2)
	go data.GOR(&wg)
	go data.GOR(&wg)
	wg.Wait()
	fmt.Println(data)
}
func (data *Embedding) GOR(wg *sync.WaitGroup) {
	defer wg.Done()
	data.mu.Lock()
	data.data++
	data.mu.Unlock()
}

//--------------------------------
//2

type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}
func COUT() {
	var wg sync.WaitGroup
	counter := &Counter{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	fmt.Println("Final count:", counter.Value())
}

/*
Keep the Mutex Unexported
You should not export the Mutex (i.e., use mu not Mu) to:
Prevent external code from accidentally or intentionally acquiring/releasing the lock.
Maintain encapsulation — external packages shouldn’t manage your struct’s synchronization.
*/

/*
Use Pointer Receivers
Use pointer receivers for methods that lock/unlock the mutex, to ensure you’re locking the same underlying Mutex.
*/

/*
Avoid Copying Structs with Mutex
If your struct contains a sync.Mutex, "do not copy it after first use" — it leads to undefined behavior.
counter := &Counter{} // use a pointer
*/
type BankAccount struct {
	mu      sync.Mutex
	balance int
}

func (a *BankAccount) Deposit(amount int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.balance += amount
}

func (a *BankAccount) Withdraw(amount int) bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.balance >= amount {
		a.balance -= amount
		return true
	}
	return false
}
