package content

/*
Garbage Collection is a form of automatic memory management.

Its job is to automatically find and reclaim(తిరిగి పొందండి) memory in a program that is no longer needed—specifically, memory that is no longer reachable or in use by the program.

Go's garbage collector (GC) manages memory within the heap.

The heap is a contiguous region of memory where dynamically allocated objects, such as structs, slices, and maps, are stored.

Unlike some languages with generational GC, Go uses a non-generational approach, meaning all objects are managed in a single, unified heap, regardless of their age.

-------------------------------------------------------------------------------------------------------
Basic Concept:
Every time you create a variable or allocate memory in a program, that memory is used to store data.

Once you no longer need that data (and the program doesn't reference it anymore), the memory becomes "garbage." Without a mechanism to reclaim it:

Your application consumes more memory.

It may eventually crash due to out-of-memory (OOM) errors.
-------------------------------------------------------------------------------------------------------
Manual vs Automatic Memory Management
| Feature                  | Manual (e.g., C, C++)    | Automatic (e.g., Go, Java)       |
| ------------------------ | ------------------------ | -------------------------------- |
| Memory Allocation        | Done manually (`malloc`) | Done automatically (`new`, etc.) |
| Memory Deallocation      | Done manually (`free`)   | Done automatically (by GC)       |
| Risk of Memory Leak      | High                     | Lower (but still possible)       |
| Risk of Dangling Pointer | High                     | Zero (Go prevents it)            |


-------------------------------------------------------------------------------------------------------

Why Does Go Need Garbage Collection?

Go is a systems-level language with a modern twist: it's designed for safety, simplicity, and scalability. These goals make GC not just useful, but essential.

Who Manages It?
The Go runtime, which is part of the standard Go distribution, includes a concurrent garbage collector that:

Tracks memory allocations (done using new, make, or implicitly during variable creation).

Identifies unreachable memory (data no longer referenced).

| Component            | Role in Garbage Collection |
| -------------------- | -------------------------- |
| **Go runtime**       | Main controller of GC      |
| **Go scheduler**     | Coordinates timing of GC   |
| **Memory manager**   | Monitors allocations       |
| **Programmer (you)** | No manual memory freeing   |

The Go runtime is a core part of the Go programming language, and it is implemented as a built-in package called runtime. It's located within the Go source tree, and it is compiled directly into your Go binary when you build your programs.

The Go runtime includes:

Garbage collector
Goroutine scheduler
Memory allocator (heap, stack management)
Panic/recover system
Channels and select implementation
System call wrappers
Time and tickers
Synchronization primitives (mutexes, semaphores, etc.)
*/

/*
Where Go Is Installed on Your System
Go has two important environment variables:
| Variable | Description                        |
| -------- | ---------------------------------- |
| `GOROOT` | Where Go itself is installed       |
| `GOPATH` | Where your Go workspace is located |

go env GOROOT
go env GOPATH
PS C:\Users\mahindra.gajjala\Desktop\GO\garbage_collector> go env GOROOT
C:\Program Files\Go
PS C:\Users\mahindra.gajjala\Desktop\GO\garbage_collector> go env GOPATH
C:\Users\mahindra.gajjala\go
PS C:\Users\mahindra.gajjala\Desktop\GO\garbage_collector>


Where the Go Runtime Is Located
The Go runtime source code is part of the standard Go installation.
$GOROOT/src/runtime/

Inside this directory:
runtime.go → Main runtime definitions
proc.go → Goroutine creation, scheduling
malloc.go → Heap allocation logic
mgc*.go → Garbage collector (mark & sweep)
panic.go → Panic and recover handling
sys_linux_amd64.s → Architecture-specific system code

This is compiled into your final binary when you run go build.

*/
/*
What Happens When You Run go run or go build?
1. Compiler Phase (cmd/compile)
Go reads your .go files
Compiles them into machine code
Includes necessary runtime code (GC, scheduler, memory management, etc.)
Stores metadata (stack size, symbols, etc.)

2. Linker Phase (cmd/link)
Combines your compiled code with:
The Go runtime
Any imported packages
Produces a standalone binary

3. Final Binary
Your compiled binary has:
Your code
All your package dependencies
The entire Go runtime (compiled in)
*/
