package directions

/*
Restricting a channel’s direction when passing to a function
Helps in API design to prevent misuse
*/
func Producer(ch chan<- int) { //sender

}
func Consumer(ch <-chan int) { //receiver

}
