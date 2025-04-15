package code

import "fmt"

func Struct_with_Different_Copies() {
	c1 := Counter{Value: 10}
	c2 := c1 // COPY created

	c2.Value = 20

	fmt.Println("c1:", c1) // original is unchanged
	fmt.Println("c2:", c2) // modified copy
}
