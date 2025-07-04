package operations

/*
There’s no built-in stop for goroutines. You must:
Use return statements
Use channels or context cancellation
*/
func Exit_stop_GoRoutine() {
	go GoRoutine_exit_stop()
}
func GoRoutine_exit_stop() {
	return
}

/*
see the context package in go

Package context defines the Context type, which carries deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes.
*/
