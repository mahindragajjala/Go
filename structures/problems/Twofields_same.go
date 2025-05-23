package problems

import (
	"fmt"
)

type Studentstr struct {
	Name string
}
type Teacherstr struct {
	Name string
}
type class struct {
	teacher Teacherstr
	Student Studentstr
}

func TwoFields_same_name() {
	school := class{teacher: Teacherstr{Name: "sai"}, Student: Studentstr{Name: "geetha"}}
	fmt.Println(school.teacher.Name)
	fmt.Println(school.Student.Name)
}
