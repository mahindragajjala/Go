package code

import (
	"fmt"
	"io"
	"strings"
)

func IO_READER() {
	// Create a new Reader from a string
	data := "Hello, Golang Reader!"
	reader := strings.NewReader(data)
	fmt.Println(reader)
	// Create a byte slice to hold the read data
	buffer := make([]byte, 8)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break // End of file reached
		}
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}
		// Print the read bytes as a string
		fmt.Printf("Read %d bytes: %s\n", n, buffer[:n])
	}
}

/*
//Data Flow & Memory Management Diagram
+---------------------------+
|      String Data          |
|  "Hello, Golang Reader!"  |
+---------------------------+
|  H  |  e  |  l  |  l  |  o  |  ,  |  ' ' |  G  |  o  |  l  |  a  |  n  |  g  |  ' ' |  R  |  e  |  a  |  d  |  e  |  r  |  !  |
|----|----|----|----|----|----|-----|----|----|----|----|----|----|-----|----|----|----|----|----|----|----|
| ^                                          |                                           |                                           |
| |---- Internal Reader Cursor ---->|                                           |                                           |
| Initial cursor                     moves on each Read()   until reaches End (EOF).

 Buffer Allocation
 +----------------+
|   buffer (8 bytes)   |
|----------------|
| []byte{ _, _, _, _, _, _, _, _ } |
+----------------+
At each loop cycle:
1️⃣ First Read:
+-------------------------------------------+
| Reads "Hello, G" into buffer              |
| buffer = [ H | e | l | l | o | , | ' ' | G ] |
| n = 8                                     |
+-------------------------------------------+

2️⃣ Second Read:
+-------------------------------------------+
| Reads "olang Re" into buffer              |
| buffer = [ o | l | a | n | g | ' ' | R | e ] |
| n = 8                                     |
+-------------------------------------------+

3️⃣ Third Read:
+-------------------------------------------+
| Reads "ader!" into buffer                 |
| buffer = [ a | d | e | r | ! ] (only 5 bytes) |
| n = 5                                     |
| err = io.EOF                              |
+-------------------------------------------+

*/
