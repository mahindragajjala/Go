package concurrentreader

/*
A concurrent reader is a goroutine that reads data in parallel with other goroutines.

The source can be a
"channel",
"shared data structure", or
"stream".

Proper synchronization ensures no two goroutines interfere.
*/
/*
Why Channels and Goroutines?
🔹 Goroutines:
Lightweight threads managed by Go runtime.
Run concurrently and can handle I/O or processing tasks.

🔹 Channels:
Provide safe communication between goroutines.

"Synchronize data transfer—only one reader and one writer at a time per channel."

When multiple goroutines read from a single channel, they automatically coordinate, meaning:

Each goroutine pulls a unique value from the channel.

There’s no need for locks.
*/
