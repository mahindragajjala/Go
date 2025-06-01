package done

/*

BEFORE READING THESE - READ EMPTY STRUCT INDETAIL
(PREREQUSITE...)




In Go, the done channel (often a chan struct{}) is a common concurrency pattern used in goroutines to signal cancellation, completion, or timeout. This is essential for writing clean, efficient, and leak-free concurrent Go programs.

🔍 What is done in Go?

The done channel is typically used to:
Signal cancellation to goroutines.
Gracefully stop long-running operations.
Avoid goroutine leaks.
Handle timeouts or context expiration.

🧠 Why use done?

Using a done channel enables:
Coordinated shutdown of multiple goroutines.
Clean resource release (e.g., closing DB connections).
Timeout-aware operations.

*/
