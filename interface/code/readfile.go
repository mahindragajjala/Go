package code

import (
	"fmt"
	"io"
	"os"
)

// printContents reads from the given io.Reader and prints its contents.
func printContents(r io.Reader) {
	data, err := io.ReadAll(r)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}
	fmt.Println("File Contents:\n", string(data))
}

func READFILE() {
	// Open a file for reading.
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Pass the file (which implements io.Reader) to the function.
	printContents(file)
}
