package mutex

import "sync/atomic"

/*
Lock-free operations on shared variables.
Extremely fast — no context switch, purely CPU instructions.
Good for counters, flags, small integers.
*/
func Atomic_operations() {
	var counter int64
	atomic.AddInt64(&counter, 1)

}
