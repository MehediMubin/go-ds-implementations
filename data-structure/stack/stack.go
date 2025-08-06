// To run this file as a standalone program, change the package name to `main` and uncomment the `main()` function.
package stack

import (
	"slices"
	"fmt"
)

type Stack struct {
	items []int
}

func (stk *Stack) Push(number int) {
	stk.items = append(stk.items, number)
}

func (stk *Stack) Pop() int {
	if stk.isEmpty() {
		fmt.Println("There is no element in the Stack")
		return 0
	}

	lastItem := stk.items[len(stk.items) - 1]
	stk.items = stk.items[:len(stk.items) - 1]
	return lastItem
}

func (stk *Stack) Top() int {
	if stk.isEmpty() {
		fmt.Println("There is no element in the Stack")
		return 0
	}

	return stk.items[len(stk.items) - 1]
}

func (stk *Stack) isEmpty() bool {
	return stk.Size() == 0
}

func (stk *Stack) Size() int {
	return len(stk.items)
}

func (stk *Stack) Clear() {
	fmt.Println("Clearing the stack....")
	stk.items = stk.items[:0]
	fmt.Println("Cleared the stack.")
}

func (stk *Stack) contains(number int) bool {
	return slices.Contains(stk.items, number)
}

func (stk *Stack) Print() {
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

// func main() {
// 	var stk Stack

// 	stk.Push(27)
// 	stk.Push(10)
// 	stk.Push(12)

// 	fmt.Println(stk.Top())
// 	fmt.Println(stk.Pop())

// 	if stk.contains(12) {
// 		fmt.Println(12, "is in the stack")
// 	} else {
// 		fmt.Println(12, "is not in the stack")
// 	}

// 	// stk.clear()

// 	stk.Print()
// }