package code

import "fmt"

/*
Takes zero bytes of memory because it has no fields, making it super lightweight.
*/
type Empty struct{}

func Empty_struct() {
	fmt.Println(Empty{})

	//Adding -  In Go, once a struct is declared, you can't add fields to it at runtime.
	/*
		If you want to "add fields" you must:
		Manually change the struct in your source code before compiling.

		Use a map or an interface{} if you want flexibility.
	*/
	person := make(map[string]interface{})

	person["Name"] = "John"
	person["Age"] = 30
	person["City"] = "Delhi"

	fmt.Println(person)

	//WE WILL USE IN -

	//As a Signal or Marker
	//Sometimes you don’t need data — you just want to signal existence.
	type Event struct{}
	fmt.Println(Event{})
	//You can use it in a map like this.
	visited := make(map[string]struct{})

	visited["apple"] = struct{}{}
	visited["banana"] = struct{}{}
	/*
		Here, you're only interested in whether the key ("apple", "banana") exists — you don't care about values.
		Using struct{} saves space because it uses zero memory compared to bool, int, or other types.
	*/

	//Implementing Sets
	//Go has no built-in Set type.
	//People use map[string]struct{} to create sets:

	set := make(map[string]struct{})

	set["A"] = struct{}{}
	set["B"] = struct{}{}

	if _, exists := set["A"]; exists {
		fmt.Println("A is in the set!")
	}

	//Signaling via Channels
	done := make(chan struct{})

	go func() {
		fmt.Println("Task started")
		done <- struct{}{} // send signal
	}()

	<-done // wait for signal
	fmt.Println("Task finished")

}

//Method Receivers Without Data
type Logger struct{}

func (Logger) Info(msg string) {
	fmt.Println("INFO:", msg)
}
