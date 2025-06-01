package codingquestions

import (
	"fmt"
	"sync"
)

/*
Bank account simulation

	Create a bank account struct with deposit and withdraw methods protected by a Mutex.
*/

type Account_details struct {
	Name        string
	TotalAmount int
	mub         sync.Mutex
}

func (account *Account_details) Deposit(credit int, wg *sync.WaitGroup) {
	defer wg.Done()
	account.mub.Lock()
	account.TotalAmount = account.TotalAmount + credit
	fmt.Println("CREDITED:", account.TotalAmount)
	account.mub.Unlock()

}
func (account *Account_details) Withdraw(withdraw int, wg *sync.WaitGroup) {
	defer wg.Done()
	account.mub.Lock()
	account.TotalAmount = account.TotalAmount - withdraw
	fmt.Println("WITHDRAWN..:", account.TotalAmount)
	account.mub.Unlock()

}
func Accounts() {

	var wgb sync.WaitGroup
	User1 := Account_details{
		Name:        "User1",
		TotalAmount: 1000,
	}
	wgb.Add(2)
	go User1.Deposit(500, &wgb)
	go User1.Withdraw(1, &wgb)
	wgb.Wait()
}

//SOLUTION - 1
/*
you're launching both operations as separate goroutines at the same time.
Due to this, the order is non-deterministic — sometimes withdrawal happens first.

Use one WaitGroup for deposit, and another for withdraw

func Accounts() {
	var wgDeposit sync.WaitGroup
	var wgWithdraw sync.WaitGroup

	User1 := Account_details{
		Name:        "User1",
		TotalAmount: 1000,
	}

	wgDeposit.Add(1)
	go User1.Deposit(500, &wgDeposit)

	// Wait for deposit to finish before withdrawal
	wgDeposit.Wait()

	wgWithdraw.Add(1)
	go User1.Withdraw(1, &wgWithdraw)
	wgWithdraw.Wait()
}
*/

//SOLUTION - 2
/*
You can also use a channel to signal that deposit is done.

func Accounts() {
	done := make(chan bool)

	User1 := Account_details{
		Name:        "User1",
		TotalAmount: 1000,
	}

	go func() {
		var wg sync.WaitGroup
		wg.Add(1)
		User1.Deposit(500, &wg)
		wg.Wait()
		done <- true // signal deposit complete
	}()

	go func() {
		<-done // wait for deposit
		var wg sync.WaitGroup
		wg.Add(1)
		User1.Withdraw(1, &wg)
		wg.Wait()
	}()

	time.Sleep(1 * time.Second) // wait for goroutines to finish
}

*/
