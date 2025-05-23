package problems

import "fmt"

//Linkedlist
/*
Create a recursive struct for a linked list using
pointers and implement a function to insert and display nodes.
*/
type Node struct {
	data int
	Next *Node
}

func Linkedlist() {
	fill := Node{data: 1, Next: &Node{data: 2, Next: &Node{data: 3}}}
	fmt.Println(fill.data)
	fmt.Println(fill.Next.data)
	fmt.Println(fill.Next.Next.data)
}
