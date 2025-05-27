package bufferedchannel

func DeadLock() {
	ch := make(chan int, 1)
	ch <- 1
	ch <- 2 // deadlock: no receiver yet

}
