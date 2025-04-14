package arraytype

import (
	"fmt"
)

var Empty_Global_array = [4]int{}

// var Max_Global_array = [1000000]int{}
var Global_array = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func Declaration_array() {
	var a = []int{}
	fmt.Println(a)
	fmt.Printf("TYPE:%T\n", a)

	//Array Initialization
	arr := [3]int{1, 2, 3}
	arr_inferred := [...]int{4, 5, 6} // size inferred
	fmt.Printf("normal array : %v \n array_inferred : %v\n", arr, arr_inferred)
	fmt.Printf("normal array type : %T \n array_inferred type: %T", arr, arr_inferred)
}

func Accessing_Array() {
	fmt.Println(Global_array[0])
	fmt.Printf("Type of the value : %T", Global_array[0])
}

//Looping over arrays

func Looping_Over_array() {
	for index, value := range Global_array {
		fmt.Printf("index : %v value : %v\n", index, value)
	}
}

func Garbage_collection_data() {
	fmt.Println("Empty_Global_array:", Empty_Global_array)
}

// func Funcdata_Max_Global_array() {
// 	fmt.Println(Max_Global_array)
// }
