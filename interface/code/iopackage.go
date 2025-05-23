package code

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// WriteToFile writes data from an io.Reader to a file
func WriteToFile(filePath string, reader io.Reader) error {
	file, err := os.Create(filePath) // create or truncate file
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	_, err = io.Copy(file, reader) // copies from reader to file (writer)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}
	return nil
}

// ReadFromFile reads data from a file and writes to an io.Writer
func ReadFromFile(filePath string, writer io.Writer) error {
	file, err := os.Open(filePath) // open file for reading
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	_, err = io.Copy(writer, file) // copies from file (reader) to writer
	if err != nil {
		return fmt.Errorf("failed to read from file: %w", err)
	}
	return nil
}

func IO_package() {
	// Sample data to write
	data := "Hello, Golang io.Reader and io.Writer Example!"

	// Convert string to io.Reader
	reader := io.NopCloser(strings.NewReader(data))

	// Write to file
	err := WriteToFile("example.txt", reader)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("File written successfully!")

	// Read from file and write to stdout
	fmt.Println("Reading file content:")
	err = ReadFromFile("example.txt", os.Stdout)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
}
