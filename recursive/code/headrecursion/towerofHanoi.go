package headrecursion

import "fmt"

func TowerOfHanoi(n int, source, destination, auxiliary string) {
	if n == 1 {
		fmt.Printf("Move disk 1 from %s to %s\n", source, destination)
		return
	}
	TowerOfHanoi(n-1, source, auxiliary, destination)
	fmt.Printf("Move disk %d from %s to %s\n", n, source, destination)
	TowerOfHanoi(n-1, auxiliary, destination, source)
}
