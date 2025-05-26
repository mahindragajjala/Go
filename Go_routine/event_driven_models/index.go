package eventdrivenmodels

/*
Event-Driven Models in Go (Golang)
Go doesn't have built-in UI or reactive libraries, but it offers concurrency primitives to create event-driven systems:

1. Channel-based Event Handling
Channels act like event buses.

Use select to react to multiple events.

select {
case msg := <-eventChannel:
    handleEvent(msg)
case <-time.After(2 * time.Second):
    fmt.Println("Timeout")
}
2. Goroutine-per-Event Source
Each event source (file, socket, DB) runs in its own goroutine and sends events through channels.

go listenForNetworkEvents(netChan)
go listenForUserInput(userChan)
3. Fan-in and Fan-out Patterns
Combine multiple event channels into one (fan-in) or distribute one event to many handlers (fan-out).

4. Context-based Cancellation for Events
Use context.Context to cancel event watchers or handlers gracefully.

5. Pub/Sub Systems
Implement a simple pub-sub engine using maps of subscribers and channels.


type Event struct {
    Name string
    Data interface{}
}
type EventBus struct {
    subscribers map[string][]chan Event
}
6. Select-based Event Loop
Build custom event loops by listening to many channels using select.
*/
