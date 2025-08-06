// To run this file as a standalone program, change the package name to `main` and uncomment the `main()` function.
package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) Prepend(data int) {
	newNode := &Node{data: data, next: list.head}
	list.head = newNode
}

func (list *LinkedList) Append(data int) {
	newNode := &Node{data: data, next: nil}

	if list.isEmpty() {
		list.head = newNode
		return
	}

	element := list.head 
	for element.next != nil {
		element = element.next
	}
	element.next = newNode
}

func (list *LinkedList) PopFront() int {
	if list.isEmpty() {
		fmt.Println("The LinkedList is empty")
		return 0
	}

	data := list.head.data
	list.head = list.head.next
	return data
}

func (list *LinkedList) PopBack() int {
	if list.isEmpty() {
		fmt.Println("The LinkedList is empty")
		return 0
	}

	if list.head.next == nil {
		data := list.head.data
		list.head = nil
		return data
	}

	element := list.head
	for element.next.next != nil {
		element = element.next
	}

	data := element.next.data
	element.next = nil
	return data
}

func (list *LinkedList) isEmpty() bool {
	return list.head == nil
}

func (list *LinkedList) Size() int {
	if list.isEmpty() {
		return 0
	}

	count := 0
	for element := list.head; element != nil; element = element.next {
		count++
	}
	return count
}

func (list *LinkedList) Contains(value int) bool {
	for element := list.head; element != nil; element = element.next {
		if element.data == value {
			return true
		}
	} 
	return false
}

func (list *LinkedList) Print() {
	if list.isEmpty() {
		fmt.Println("The LinkedList is empty")
		return
	}

	fmt.Println("Printing the full linkedlist...")
	for element := list.head; element != nil ; element = element.next {
		fmt.Println(element.data)
	}
}

func (list *LinkedList) Clear() {
	list.head = nil
}

// func main() {
// 	var list LinkedList

// 	list.Append(10)
// 	list.Append(20)
// 	list.Append(30)

// 	list.PopBack()
// 	list.PopFront()

// 	fmt.Println("Size of list:", list.Size())

// 	value := 10
// 	if list.Contains(value) {
// 		fmt.Println(value, "exists")
// 	} else {
// 		fmt.Println(value, "doesn't exist")
// 	}

// 	list.Print()
// }