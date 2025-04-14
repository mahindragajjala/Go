package structures

import "fmt"

type Student struct {
	Name       string
	RollNumber int
	TotalMarks int
}

func Structures_Go() {
	//std1 - declaration / Student{} - initialization
	std1 := Student{
		Name:       "tommy",
		RollNumber: 222,
		TotalMarks: 38,
	}
	_ = std1

	students := []Student{
		{Name: "Alice", RollNumber: 1, TotalMarks: 90},
		{Name: "Bob", RollNumber: 2, TotalMarks: 85},
	}
	students[1].Name = "Slice"
	fmt.Println("students--->", students)
}
