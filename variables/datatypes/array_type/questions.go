package arraytype

import "fmt"

var data = []int{56, 24, 78, 45, 2, 455, 777}

//brute/better/optimal
func Largest_element_in_a_array() {
	var largest_element int
	for i := 0; i < len(data); i++ {
		if largest_element < data[i] {
			largest_element = data[i]
		}
	}
	fmt.Println(largest_element)
}

func New_array_creation() {
	var temp int
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i] > data[j] {
				temp = data[j]
				data[j] = data[i]
				data[i] = temp
			}
		}
	}
	fmt.Println("Sorted array in ascending order:", data)
}

func Swapping() {
	smalldata := []int{23, 4, 88}
	/*
	   23
	   4
	   88

	   4
	   88

	   88
	*/
	Temp := 0
	for i := 0; i < len(smalldata); i++ {
		fmt.Println(smalldata[i])
		for j := i + 1; j < len(smalldata); j++ {
			fmt.Println(smalldata[j])
			if data[i] < data[j] {
				// Temp = data[i]
				// data[i] = data[j]
				// data[j] = data[i]
				Temp = data[j]
				data[j] = data[i]
				data[i] = Temp
			}
		}
	}
	fmt.Println(Temp, data)
}
