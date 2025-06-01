package operatorspunctuationdata

import "fmt"

/*

In Go, the ; (semicolon) is mostly invisible to you as a programmer
— it's inserted automatically by the Go compiler (known as semicolon inference),
so you rarely need to type it yourself.

*/

func Semicolon() {
	//Writing multiple statements on a single line
	fmt.Print("A")
	fmt.Print("B")

	//Inside for loops (like C-style loops)
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

}

/*

❌ Don't do this:

x := 10;
y := 20;

*/
