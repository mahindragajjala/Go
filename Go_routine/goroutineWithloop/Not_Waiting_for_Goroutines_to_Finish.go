package goroutinewithloop

//Not Waiting for Goroutines to Finish
/*
Problem: The main function exits before goroutines complete.
Fix: Use sync.WaitGroup to wait for all goroutines.

var wg sync.WaitGroup

for i := 0; i < 5; i++ {
    wg.Add(1)
    go func(n int) {
        defer wg.Done()
        fmt.Println(n)
    }(i)
}

wg.Wait()
*/
