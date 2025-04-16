package code

import (
	"fmt"
)

// Reader Interface
type Reader interface {
	Read(p []byte) (n int, err error)
}

// Writer Interface
type Writer interface {
	Write(p []byte) (n int, err error)
}

// ReadWriter Interface embedding Reader and Writer
type ReadWriter interface {
	Reader
	Writer
}

// Struct implementing both Reader and Writer
type ReaderWriter struct {
	Name string
	Time int
}

// Correct Read method
func (a ReaderWriter) Read(p []byte) (n int, err error) {
	fmt.Println("READER NAME:", a.Name)
	return a.Time, nil
}

// Correct Write method
func (a ReaderWriter) Write(p []byte) (n int, err error) {
	fmt.Println("WRITE NAME:", a.Name)
	return a.Time, nil
}

// Function using ReadWriter interface
func Interface_function(a ReadWriter) {
	a.Read(make([]byte, 10))  // Passing dummy byte slice
	a.Write(make([]byte, 10)) // Passing dummy byte slice
}

// Main function to test
func Main_RW() {
	a := ReaderWriter{Name: "sai", Time: 21}
	Interface_function(a)
}
