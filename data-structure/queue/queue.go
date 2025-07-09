package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Queue struct {
	items []Person
}

func (q *Queue) push(name string, age int) {
	q.items = append(q.items, Person{name, age})
}

func (q *Queue) isEmpty() bool {
	if len(q.items) == 0 {
		return true
	} else {
		return false
	}
}

func (q *Queue) pop() Person {
	if q.isEmpty() {
		fmt.Println("There is no element in the Queue")
	}

	firstElem := q.items[0]
	q.items = q.items[1:]
	return firstElem
}

func (q *Queue) front() Person {
	if q.isEmpty() {
		fmt.Println("There is no element in the Queue")
	}

	return q.items[0]
}

func (q *Queue) back() Person {
	if q.isEmpty() {
		fmt.Println("There is no element in the Queue")
	}

	return q.items[len(q.items) - 1]
}

func (q *Queue) size() int {
	return len(q.items)
}

func (q *Queue) clear() {
	q.items = q.items[:0]
}

func (q *Queue) contains(name string, age int) bool {
	for _, person := range q.items {
		if person.name == name && person.age == age {
			return true
		}
	}
	return false
}

func main() {
	var q Queue
	q.push("Mehedi", 10)
	q.push("Hasan", 11)
	q.push("Mubin", 12)
	q.push("Muhin", 13)

	fmt.Println("First element in the Queue", q.front())
	fmt.Println("Last element in the Queue", q.back())
	fmt.Println("Size of the Queue: ", q.size())

	fmt.Println("Element Popped", q.pop())		
	fmt.Println("Current First element in the Queue:", q.front())
	
	q.clear()
	fmt.Println("Cleared the Queue")
	fmt.Println("Current size of the Queue:", q.size())
}