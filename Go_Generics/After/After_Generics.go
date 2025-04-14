package after

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func SumValues[K comparable, V int | float64](m map[K]V) V {
	var total V
	for _, v := range m {
		total += v
	}
	return total
}

func After_generics_main() {
	m := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5}
	l := SumValues(m)
	fmt.Println(l)
	n := map[string]float64{"A": 1.2, "B": 2.2, "C": 3.3, "D": 4.2, "E": 5.4}
	v := SumValues(n)
	fmt.Println(v)

	//app := map[string]float64{"A": 1.2, "B": 2.2, "C": 3.3, "D": 4.2, "E": 5.4}
	Manual(5, 2)
	Manual(5.1, 2.1)

}

func Manual[A int | float64 | float32](a A, b A) A {
	sum := a + b
	return sum
}

type Data interface {
	int | float64 | float32
}

func Manual_Structure[a Data](a) {

}

// using package
func Using_package[i constraints.Ordered](i){
	fmt.Println(i)
}
