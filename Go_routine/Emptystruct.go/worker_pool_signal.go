package emptystructgo

/*
done := make(chan struct{})

go func() {
    for {
        select {
        case <-done:
            fmt.Println("worker stopped")
            return
        default:
            fmt.Println("working...")
            time.Sleep(500 * time.Millisecond)
        }
    }
}()

time.Sleep(2 * time.Second)
close(done) // stop the worker

*/
