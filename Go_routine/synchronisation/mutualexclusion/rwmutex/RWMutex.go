package rwmutex

import (
	"fmt"
	"sync"
)

/*
----------------------------Allows multiple readers but only one writer------------------------

RLock() / RUnlock() — Shared read lock
Lock() / Unlock() — Exclusive write lock

Readers only read data and writers modify it.
When writer has access to data then none other thread (reader or writer) can share such access.

#1
Imagine a writer is updating a big document, but they can only make changes bit by bit — not all at once. While the writer is still working, the document is incomplete and messy.

Now, if a reader tries to read the document while the writer is still updating it, the reader might see some parts changed and some parts not yet changed. This could confuse the reader or cause mistakes.

So, to keep things safe and correct, we block the reader from seeing the document until the writer finishes all the changes completely. That way, the reader only ever sees a fully finished, clean version — no broken or half-updated data.

#2
When many people (or threads) want to read or write data, we want to be fair and make sure no one gets stuck waiting forever. The different rules are about preventing starvation, which means waiting endlessly without getting a chance to do your work.

Writers cannot be starved:
Writers (people who want to make changes) should not have to wait forever. They must get their turn to write at some point, even if many readers are trying to read.

Readers cannot be starved:
Readers (people who want to look at the data) should also not have to wait forever. They must be allowed to read eventually, even if many writers want to write.

No thread shall be allowed to starve:
This means everyone — whether a reader or a writer — must get a fair chance. No one should be stuck waiting forever while others keep going ahead.
*/

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
RLock
For the sake of brevity let’s skip parts related to race detector (they will be replaced by ...).

func (rw *RWMutex) RLock() {
    ...
    if atomic.AddInt32(&rw.readerCount, 1) < 0 {
        runtime_SemacquireMutex(&rw.readerSem, false)
    }
    ...
}
Field readerCount is an int32 value indicating number of pending readers — already reading data or blocked by writer.
This is basically the number of readers which have called RLock function but no RUnlock yet.

atomic.AddInt32 is the atomic equivalent of:

*addr += delta
return *addr
where addr has type *int32 and delta has type int32. Because it’s atomic operation there is no risk that adding delta will interfere with other threads (more about fetch-and-add).

If writers don’t come into play then readerCount is always greater than or equal to 0 and readers have fast non-blocking path involving only calls of atomic.AddInt32.
*/
