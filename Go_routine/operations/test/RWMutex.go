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
