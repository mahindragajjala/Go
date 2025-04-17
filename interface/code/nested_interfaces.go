package code

import "fmt"

// Smaller Interfaces
type Readerinterface interface {
	Read(p []byte) (n int, err error)
}

type Writerinterface interface {
	Write(p []byte) (n int, err error)
}

// Composed Interface
type ReadWriterinterface interface {
	Readerinterface
	Writerinterface
}

// Struct implementing both Reader and Writer
type File struct{}

func (f File) Read(p []byte) (int, error) {
	fmt.Println("Reading data...")
	return len(p), nil
}

func (f File) Write(p []byte) (int, error) {
	fmt.Println("Writing data...")
	return len(p), nil
}

func Nested_interface() {
	var rw ReadWriterinterface = File{}
	rw.Read(make([]byte, 10))
	rw.Write([]byte("Hello"))
}
