package mutex

import (
	"fmt"
	"sync"
)

/*
Real-World Example: Bank Account System

Imagine you have a backend service managing bank accounts. There are two operations:

Withdraw(amount int) – Reduces balance.
ValidateTransaction() – Validates internal rules before allowing the transaction.
Now suppose Withdraw() calls ValidateTransaction() internally, but both functions try to lock the same sync.Mutex to protect shared balance data.
*/

//DEAD LOCK HAPPENS - DUE TO DOUBLE DEAD LOCK...
/*
package main

import (
    "fmt"
    "sync"
)

type Account struct {
    balance int
    mu      sync.Mutex
}

func (a *Account) Withdraw(amount int) {
    a.mu.Lock()
    defer a.mu.Unlock()

    fmt.Println("Withdraw: Acquired lock")

    if a.ValidateTransaction() {
        a.balance -= amount
        fmt.Println("Withdraw: Success, new balance:", a.balance)
    } else {
        fmt.Println("Withdraw: Validation failed")
    }
}

func (a *Account) ValidateTransaction() bool {
    a.mu.Lock()         // ❌ Deadlock here — trying to lock again
    defer a.mu.Unlock()

    fmt.Println("ValidateTransaction: Accessing account safely")
    return a.balance > 0
}

func main() {
    acc := &Account{balance: 100}
    acc.Withdraw(50) // DEADLOCK occurs
}

What Happens Here?

Withdraw() locks the mutex to ensure exclusive access to balance.
It calls ValidateTransaction(), which again tries to lock the mutex.
Since Go's sync.Mutex is not reentrant, it blocks forever waiting for a lock it already holds
*/

//SOLUTION-

// Option 1: Split Lock Responsibility
type Account struct {
	balance int
	mu      sync.Mutex
}

func (a *Account) Withdraw(amount int) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.balance > 0 {
		a.balance -= amount
		fmt.Println("Withdraw: Success, new balance:", a.balance)
	} else {
		fmt.Println("Withdraw: Insufficient balance")
	}
}

// No need for a mutex inside ValidateTransaction
func (a *Account) ValidateTransaction() bool {
	// Assume balance was already locked outside
	return a.balance > 0
}

/*
Here, Withdraw() is fully responsible for locking. Internal methods assume the data is already protected.
*/
/*
Imagine a security vault room where only one key is allowed at a time.
The bank manager (goroutine) enters and locks the room (mutex).
While inside, they try to use the same key again to open another secure cabinet — but the door doesn't allow the same key to be reused from the inside.

This causes them to stand forever holding the key, never finishing their task — exactly like a deadlock in a non-reentrant mutex.
*/
