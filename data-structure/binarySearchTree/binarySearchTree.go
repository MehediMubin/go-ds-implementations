// To run this file as a standalone program, change the package name to `main` and uncomment the `main()` function.
package binarySearchTree

import "fmt"

type Node struct {
	Value int
	Left *Node 
	Right *Node
}

func (n *Node) Insert(value int) *Node {
	if n == nil {
		return &Node{Value: value}
	}

	if value < n.Value {
		n.Left = n.Left.Insert(value)
	} else if value > n.Value {
		n.Right = n.Right.Insert(value)
	}

	return n
}

func (n *Node) Search(value int) bool {
	if n == nil {
		return false
	}

	if value == n.Value {
		return true
	} else if (value < n.Value) {
		return n.Left.Search(value)
	} else {
		return n.Right.Search(value)
	}
}

func (n *Node) InOrder() {
	if n == nil {
		return
	}

	n.Left.InOrder()
	fmt.Print(n.Value, " ")
	n.Right.InOrder()
}

func (n *Node) PreOrder() {
	if n == nil {
		return
	}

	fmt.Print(n.Value, " ")
	n.Left.PreOrder()
	n.Right.PreOrder()
}

func (n *Node) PostOrder() {
	if n == nil {
		return
	}

	n.Left.PostOrder()
	n.Right.PostOrder()
	fmt.Print(n.Value, " ")
}

// func main() {
// 	var root *Node

// 	values := []int{3, 1, 2, 5, 4, 6}
// 	for _, v := range values {
// 		root = root.Insert(v)
// 	}

// 	root.PostOrder()
// 	fmt.Println()
// }