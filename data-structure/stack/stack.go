package main

import (
	"slices"
	"fmt"
)

type Person struct {
	name string
	age  int
}

type Stack struct {
	items []Person
}

func (stk *Stack) push(p Person) {
	stk.items = append(stk.items, p)
}

func (stk *Stack) pop() Person {
	if stk.isEmpty() {
		fmt.Println("There is no element in the Stack")
		return Person{}
	}

	lastItem := stk.items[len(stk.items) - 1]
	stk.items = stk.items[:len(stk.items) - 1]
	return lastItem
}

func (stk *Stack) top() Person {
	if stk.isEmpty() {
		fmt.Println("There is no element in the Stack")
		return Person{}
	}

	return stk.items[len(stk.items) - 1]
}

func (stk *Stack) isEmpty() bool {
	return stk.size() == 0
}

func (stk *Stack) size() int {
	return len(stk.items)
}

func (stk *Stack) clear() {
	fmt.Println("Clearing the stack....")
	stk.items = stk.items[:0]
	fmt.Println("Cleared the stack.")
}

func (stk *Stack) contains(item Person) bool {
	return slices.Contains(stk.items, item)
}

func (stk *Stack) print() {
	if stk.isEmpty() {
		fmt.Println("There is no item in the Stack")
		return
	}

	fmt.Println("Printing items in the Stack...")
	for i := len(stk.items) - 1; i >= 0; i-- {
		fmt.Println(stk.items[i])
	}
	fmt.Println("Printing Done")
}

func main() {
	var stk Stack

	stk.push(Person{"Mehedi", 27})
	stk.push(Person{"Hasan", 10})
	stk.push(Person{"Mubin", 12})

	fmt.Println(stk.top())
	fmt.Println(stk.pop())

	if stk.contains(Person{"Mubin", 12}) {
		fmt.Println("Mubin is in the stack")
	} else {
		fmt.Println("Mubin is not in the stack")
	}

	// stk.clear()

	stk.print()
}