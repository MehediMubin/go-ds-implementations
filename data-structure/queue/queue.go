package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Push(number int) {
	q.items = append(q.items, number)
}

func (q *Queue) isEmpty() bool {
	if len(q.items) == 0 {
		return true
	} else {
		return false
	}
}

func (q *Queue) Pop() int {
	if q.isEmpty() {
		fmt.Println("There is no element in the Queue")
	}

	firstElem := q.items[0]
	q.items = q.items[1:]
	return firstElem
}

func (q *Queue) Front() int {
	if q.isEmpty() {
		fmt.Println("There is no element in the Queue")
	}

	return q.items[0]
}

func (q *Queue) Back() int {
	if q.isEmpty() {
		fmt.Println("There is no element in the Queue")
	}

	return q.items[len(q.items) - 1]
}

func (q *Queue) Size() int {
	return len(q.items)
}

func (q *Queue) Clear() {
	q.items = q.items[:0]
}

func (q *Queue) Contains(number int) bool {
	for _, value := range q.items {
		if value == number {
			return true
		}
	}
	return false
}

func main() {
	var q Queue
	q.Push(10)
	q.Push(11)
	q.Push(12)
	q.Push(13)

	fmt.Println("First element in the Queue", q.Front())
	fmt.Println("Last element in the Queue", q.Back())
	fmt.Println("Size of the Queue: ", q.Size())

	fmt.Println("Element Popped", q.Pop())		
	fmt.Println("Current First element in the Queue:", q.Front())
	
	q.Clear()
	fmt.Println("Cleared the Queue")
	fmt.Println("Current size of the Queue:", q.Size())
}