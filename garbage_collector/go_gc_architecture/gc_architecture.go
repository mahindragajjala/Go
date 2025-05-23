package gogcarchitecture

/*
Go's GC Architecture
Mark-and-sweep algorithm
Tricolor abstraction (white, grey, black objects)
Non-generational vs generational GC


Go uses a concurrent, tri-color mark-and-sweep garbage collector.
It is non-generational (as of Go 1.22), and is designed to work with low pause times, making it suitable for server applications.

⚙️ High-level Phases of Go GC
Start GC Cycle (STW - Stop the World briefly)

Concurrent Mark Phase – runs alongside your application (mutator).
Concurrent Sweep Phase
Final STW – very short pause to finalize sweeping.

Go's garbage collector (GC) employs a concurrent tri-color mark-and-sweep algorithm, which is designed to minimize stop-the-world pauses and optimize performance for concurrent programs.

It's a non-generational GC, meaning all objects are managed in a single heap, regardless of their lifespan. The GC runs concurrently with the application, using write barriers to track object modifications and ensure that memory is reclaimed efficiently.

*/
