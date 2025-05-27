package selectkeyword

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

/*
Using `select` with `context.Context` (Cancellation)

* Graceful goroutine shutdown.
* Avoiding goroutine leaks.
* Listening for cancellation signals.
*/
/*
Using select with "context.Context" is a powerful idiom in Go for writing cancellation-aware goroutines

A cancellation-aware goroutine is a goroutine that:
Listens for a signal to stop (usually using context.Context) and exits gracefully when told to do so.
This is important because Go goroutines don't automatically stop — you must tell them when to stop

Think of a goroutine like a worker in a factory.
If you don't tell the worker to stop (even when the shift ends), they keep working forever. 😓
A cancellation-aware worker checks a clock (like ctx.Done()) to know when the shift is over and then cleans up and leaves. ✅

which helps with:
✅ Graceful goroutine shutdown
✅ Avoiding goroutine leaks
✅ Listening for cancellation signals

*/
/*
Package context defines the Context type, which carries deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes.

Context provides a mechanism to control the lifecycle, cancellation, and propagation of requests across multiple goroutines.

Consider a scenario where you need to fetch data from multiple APIs concurrently. By using context, you can ensure that all the API requests are canceled if any of them exceeds a specified timeout.
*/
func Context() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	urls := []string{
		"https://api.example.com/users",
		"https://api.example.com/products",
		"https://api.example.com/orders",
	}

	results := make(chan string)

	for _, url := range urls {
		go fetchAPI(ctx, url, results)
	}

	for range urls {
		fmt.Println(<-results)
	}
}

func fetchAPI(ctx context.Context, url string, results chan<- string) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		results <- fmt.Sprintf("Error creating request for %s: %s", url, err.Error())
		return
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		results <- fmt.Sprintf("Error making request to %s: %s", url, err.Error())
		return
	}
	defer resp.Body.Close()

	results <- fmt.Sprintf("Response from %s: %d", url, resp.StatusCode)
}
