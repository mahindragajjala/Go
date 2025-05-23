package mutex

import (
	"fmt"
	"sync"
	"time"
)

/*
Printing Documents on a Shared Printer
Imagine:

There is one printer shared by several employees.
Employees want to print their documents.
If the printer is free, the employee uses it.
If the printer is busy, the employee does NOT wait but decides to try printing later or do something

*/

/*
In Go 1.18, the sync package introduced the TryLock() method for sync.Mutex, allowing you to attempt a lock without blocking.

This is useful when you want to try acquiring a lock only if it’s immediately available, rather than waiting (blocking) until it becomes available.
*/
var muT sync.Mutex

func tryWorker(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	if muT.TryLock() {
		fmt.Println(name, "acquired the lock")
		time.Sleep(2 * time.Second)
		mu.Unlock()
		fmt.Println(name, "released the lock")
	} else {
		fmt.Println(name, "could NOT acquire the lock, skipping")
	}
}

func Trylock_function() {
	var wg sync.WaitGroup
	wg.Add(2)

	go tryWorker("Worker1", &wg)
	time.Sleep(100 * time.Millisecond) // ensure Worker1 tries first
	go tryWorker("Worker2", &wg)

	wg.Wait()
}

var printer sync.Mutex

func employee(name string) {
	if printer.TryLock() {
		fmt.Println(name, "is printing a document...")
		time.Sleep(2 * time.Second) // Printing takes some time
		printer.Unlock()
		fmt.Println(name, "finished printing and freed the printer")
	} else {
		fmt.Println(name, "found the printer busy and will try later")
	}
}

func Printer_func() {
	go employee("Alice")
	time.Sleep(100 * time.Millisecond) // Stagger the attempts a bit
	go employee("Bob")
	time.Sleep(100 * time.Millisecond)
	go employee("Charlie")

	time.Sleep(5 * time.Second) // Wait for all goroutines to finish
}
