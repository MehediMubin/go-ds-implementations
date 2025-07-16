package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Node struct {
	data Person
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) prepend(data Person) {
	newNode := &Node{data: data, next: list.head}
	list.head = newNode
}

func (list *LinkedList) append(data Person) {
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

func (list *LinkedList) popFront() Person {
	if list.isEmpty() {
		fmt.Println("The LinkedList is empty")
		return Person{}
	}

	data := list.head.data
	list.head = list.head.next
	return data
}

func (list *LinkedList) popBack() Person {
	if list.isEmpty() {
		fmt.Println("The LinkedList is empty")
		return Person{}
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

func (list *LinkedList) size() int {
	if list.isEmpty() {
		return 0
	}

	count := 0
	for element := list.head; element != nil; element = element.next {
		count++
	}
	return count
}

func (list *LinkedList) contains(value Person) bool {
	for element := list.head; element != nil; element = element.next {
		if element.data == value {
			return true
		}
	} 
	return false
}

func (list *LinkedList) print() {
	if list.isEmpty() {
		fmt.Println("The LinkedList is empty")
		return
	}

	fmt.Println("Printing the full linkedlist...")
	for element := list.head; element != nil ; element = element.next {
		fmt.Println(element.data)
	}
}

func (list *LinkedList) clear() {
	list.head = nil
}

func main() {
	var list LinkedList

	list.append(Person{"Mehedi", 27})
	list.append(Person{"Hasan", 17})
	list.append(Person{"Mubin", 67})

	list.popBack()
	list.popFront()

	fmt.Println("Size of list:", list.size())

	value := Person{"Hasan", 17}
	if list.contains(value) {
		fmt.Println(value, "exists")
	} else {
		fmt.Println(value, "doesn't exist")
	}

	list.print()
}