package code

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// writeData writes a given string to any io.Writer.
func writeData(w io.Writer, data string) {
	_, err := w.Write([]byte(data))
	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}
	fmt.Println("Data written successfully!")
}

func Writefile() {
	// Example 1: Writing to a file.
	file, err := os.Create("output.txt") // Creates or overwrites the file
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writeData(file, "Hello, this is written to a file!\n")

	// Example 2: Writing to a bytes.Buffer.
	var buffer bytes.Buffer
	writeData(&buffer, "Hello, this is written to a buffer!\n")

	// Print the buffer contents.
	fmt.Println("Buffer Contents:\n", buffer.String())
}
