package questions

import (
	"fmt"
	"sort"
)

// Define the Employee struct
type Employee struct {
	Name   string
	Salary float64
}

// Create a type alias for a slice of Employee
type BySalary []Employee

// Implement sort.Interface methods for BySalary

// Len returns the number of elements in the collection
func (e BySalary) Len() int {
	return len(e)
}

// Less compares two elements and returns true if the first is less than the second
func (e BySalary) Less(i, j int) bool {
	return e[i].Salary < e[j].Salary
}

// Swap swaps the elements with indexes i and j
func (e BySalary) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func Sorting_empolyees() {
	employees := []Employee{
		{Name: "Alice", Salary: 50000},
		{Name: "Bob", Salary: 45000},
		{Name: "Charlie", Salary: 70000},
		{Name: "Dave", Salary: 55000},
	}
	fmt.Println(employees)
	fmt.Println("Before sorting:")
	for _, emp := range employees {
		fmt.Printf("%s: %.2f\n", emp.Name, emp.Salary)
	}

	// Sort employees by salary
	sort.Sort(BySalary(employees))

	fmt.Println("\nAfter sorting by salary:")
	for _, emp := range employees {
		fmt.Printf("%s: %.2f\n", emp.Name, emp.Salary)
	}
}
