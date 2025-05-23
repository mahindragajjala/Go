package referencedereference

/*
when you define recursive structs (i.e., structs that reference themselves),
there are some important rules to avoid infinite size at compile-time.

type Node struct {
    data int
    next Node  // ❌ Invalid: This will cause infinite size
}
*/
