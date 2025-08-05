package main

import "fmt"

type MinHeap struct {
	data []int
}

func (h *MinHeap) Insert(val int) {
	h.data = append(h.data, val)
	h.HeapifyUp(len(h.data) - 1)
}

func (h *MinHeap) HeapifyUp(index int) {

}

func main() {
	heap := &MinHeap{}

	heap.Insert(10)
	heap.Insert(5)
	heap.Insert(2)
	heap.Insert(3)

	fmt.Println(heap.data)
}