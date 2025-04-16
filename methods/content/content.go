package content

/*
In Go, a method is just like a function, but it's tied to a type (usually a struct).
func (receiver Type) MethodName(parameters) returnType {
    // method body
}

The receiver allows the method to act like it “belongs” to a type, similar to classes in other languages.

Memory Management of Methods
When you define a method, the behavior depends on the "receiver type":

Receiver Type	Memory Behavior
(t MyType)		Passes by value — creates a copy of the data.
(t *MyType)		Passes by reference — works on the same object in memory.


*/
