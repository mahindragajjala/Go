package problems

import "fmt"

// Tree struct with Left and Right child pointers.
type Tree struct {
	Data  int
	Left  *Tree
	Right *Tree
}

// Insert inserts a new value into the binary search tree.
func (t *Tree) Insert(value int) *Tree {
	if t == nil {
		return &Tree{Data: value}
	}
	if value < t.Data {
		t.Left = t.Left.Insert(value)
	} else {
		t.Right = t.Right.Insert(value)
	}
	return t
}

// Search looks for a value in the tree and returns true if found.
func (t *Tree) Search(value int) bool {
	if t == nil {
		return false
	}
	if t.Data == value {
		return true
	} else if value < t.Data {
		return t.Left.Search(value)
	} else {
		return t.Right.Search(value)
	}
}

// InOrder prints the tree nodes in in-order traversal.
func (t *Tree) InOrder() {
	if t == nil {
		return
	}
	t.Left.InOrder()
	fmt.Print(t.Data, " ")
	t.Right.InOrder()
}

// Example usage.
func Insertion_Tree() {
	var root *Tree

	// Insert nodes
	for i := 0; i < 5; i++ {
		root = root.Insert(i)
	}

	// Print in-order traversal
	fmt.Print("In-Order Traversal: ")
	root.InOrder()
	fmt.Println()

	// Search for a value
	valueToSearch := 3
	if root.Search(valueToSearch) {
		fmt.Printf("%d found in the tree!\n", valueToSearch)
	} else {
		fmt.Printf("%d not found in the tree.\n", valueToSearch)
	}
}
