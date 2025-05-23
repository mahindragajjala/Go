package code

import (
	"fmt"
	"io"
	"strings"
)

func Using_structure_from_string() {
	data := "these is the string from the local function"

	reader := strings.NewReader(data)
	buf := make([]byte, 10)
	reader.Read(buf)

	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Read error:", err)
			return
		}
		fmt.Printf("Read %d bytes: %s\n", n, buf[:n])
	}

}
