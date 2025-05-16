package test

import (
	"fmt"
	"sync"
)

var (
	rwMu sync.RWMutex
	data = make(map[string]string)
)

func read(key string, wg *sync.WaitGroup) {
	defer wg.Done()
	rwMu.RLock()
	fmt.Println("Reading:", key, "->", data[key])
	rwMu.RUnlock()
}

func write(key, value string, wg *sync.WaitGroup) {
	defer wg.Done()
	rwMu.Lock()
	data[key] = value
	rwMu.Unlock()
}

func RWMutex() {
	var wg sync.WaitGroup

	wg.Add(3)
	go write("name", "GoLang", &wg)
	go read("name", &wg)
	go read("name", &wg)

	wg.Wait()
}

/*
Reads are frequent, writes are rare.
You want concurrent reads and exclusive writes.
*/
/*
RWMutex stands for Read-Write Mutex.
It is a special kind of lock that allows:
Multiple readers to read at the same time,
But only one writer can write, and no readers allowed during writing.

It’s used when you want to optimize read-heavy code, where reads are safe but writes need exclusive access.

| Method      | Purpose                               |
| ----------- | ------------------------------------- |
| `RLock()`   | Acquire read lock (shared access)     |
| `RUnlock()` | Release read lock                     |
| `Lock()`    | Acquire write lock (exclusive access) |
| `Unlock()`  | Release write lock                    |


Rlocker()
Trylock()
TryRLock()
RLocker().Lock
RLocker().unlock
*/
/*
Imagine a library:

Many students can read books together — no problem.
But if someone wants to rearrange the books (write), everyone must leave the library until the rearrangement is done.
This is exactly how RWMutex works.
*/

/* var rwMu sync.RWMutex
var data = 0

func read(id int) {
	rwMu.RLock() // Acquire read lock
	fmt.Println("Reader", id, "reading:", data)
	time.Sleep(time.Millisecond * 500)
	rwMu.RUnlock() // Release read lock
}

func write(val int) {
	rwMu.Lock() // Acquire write lock
	fmt.Println("Writer writing:", val)
	data = val
	time.Sleep(time.Second)
	rwMu.Unlock() // Release write lock
}

func READWRITE_MUTEX() {
	go read(1)
	go read(2)
	go write(10)
	go read(3)

	time.Sleep(3 * time.Second)
}
*/

func RWMUTEC_function() {

}
