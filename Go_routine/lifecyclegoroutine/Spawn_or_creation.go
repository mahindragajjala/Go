package lifecyclegoroutine

/*
Spawn (Creation)
This is when a goroutine is started using the go keyword.
How it works:
You prefix a function call with go.
This tells Go to run the function in the background as a separate lightweight thread (goroutine).

go sayHello()

When this line runs, Go doesn't wait for sayHello() to finish — it just starts it and keeps going.
*/
