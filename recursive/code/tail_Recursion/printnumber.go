package tailrecursion

import "fmt"

func PrintNto1(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	PrintNto1(n - 1)
}
