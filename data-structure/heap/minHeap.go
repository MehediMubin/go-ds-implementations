package main

import "fmt"

type MinHeap struct {
	data []int
}

func (h *MinHeap) Insert(val int) {
	h.data = append(h.data, val)
	h.heapifyUp(len(h.data) - 1)
}

func (h *MinHeap) heapifyUp(index int) {
	for h.data[parent(index)] > h.data[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func (h *MinHeap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func main() {
	heap := &MinHeap{}

	heap.Insert(10)
	heap.Insert(5)
	heap.Insert(2)
	heap.Insert(3)

	fmt.Println(heap.data)
}