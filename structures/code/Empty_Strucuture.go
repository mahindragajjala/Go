package code

import "fmt"

//As a Set (map key with no value)
func Map_with_no_values() {
	mySet := make(map[string]struct{})

	mySet["apple"] = struct{}{}
	mySet["banana"] = struct{}{}

	if _, exists := mySet["apple"]; exists {
		fmt.Println("apple is in the set")
	}
}

//Channel signaling (without data)
func Channel_singnaling_without_Data() {
	done := make(chan struct{})

	go func() {
		// do some work
		done <- struct{}{} // signal completion
	}()

	<-done // wait for signal

}

// Zero-size markers in struct tags or APIs
func Zero_size_markers() {
	//You might see it in advanced patterns like embedding for marker traits:

	type NoCopy struct{}

}
