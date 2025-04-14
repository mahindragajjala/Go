package before

import "fmt"

func sumInts(m map[string]int) int {
	total := 0
	for _, v := range m {
		total += v
	}
	return total
}
func Before_Generics_main() {

	//map
	m := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5}
	data := sumInts(m)
	fmt.Println(data)
}
