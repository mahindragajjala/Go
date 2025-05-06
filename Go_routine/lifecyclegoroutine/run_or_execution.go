package lifecyclegoroutine

/*
Run (Execution)
This is when the goroutine is actually doing its work (i.e., the function is running).

Key points:

Goroutines are scheduled by the "Go runtime" (like a manager giving time to each worker).
They run concurrently (not necessarily in parallel unless you have multiple CPU cores).
The main program doesn’t control when exactly a goroutine runs — the Go scheduler handles that.

who handles the main function:
The Go runtime handles the execution of main() just like it does for goroutines.
The main() function itself runs inside a special goroutine, called the main goroutine.
This means even main() is scheduled and managed by the Go scheduler.
*/
