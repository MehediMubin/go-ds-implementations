// To run this file as a standalone program, change the package name to `main` and uncomment the `main()` function.
package maxHeap

import (
	"fmt"
)

type MaxHeap struct {
	data []int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

func (h *MaxHeap) Insert(val int) {
	h.data = append(h.data, val)
	h.heapifyUp(len(h.data) - 1)
}

func (h *MaxHeap) heapifyUp(index int) {
	for h.data[parent(index)] < h.data[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MaxHeap) Peek() (int, error) {
	if len(h.data) == 0 {
		return 0, fmt.Errorf("heap is empty")
	}

	return h.data[0], nil
}

func (h *MaxHeap) ExtractMax() (int, error) {
	if len(h.data) == 0 {
		return 0, fmt.Errorf("heap is empty")
	}

	minElement := h.data[0]
	lastIndex := len(h.data) - 1

	h.data[0] = h.data[lastIndex]
	h.data = h.data[:lastIndex]

	h.heapifyDown(0)
	return minElement, nil
}

func (h *MaxHeap) heapifyDown(index int) {
	lastIndex := len(h.data) - 1

	for {
		left, right := leftChild(index), rightChild(index)
		smallest := index 

		if left <= lastIndex && h.data[left] > h.data[smallest] {
			smallest = left
		}
		if right <= lastIndex && h.data[right] > h.data[smallest] {
			smallest = right
		}
		if smallest == index {
			break
		}

		h.swap(smallest, index)
		index = smallest
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func leftChild(i int) int {
	return 2 * i + 1
}

func rightChild(i int) int {
	return 2 * i + 2
}

func (h *MaxHeap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *MaxHeap) IsEmpty() bool {
	if len(h.data) == 0 {
		return true
	} else {
		return false
	}
}

// func main() {
// 	heap := NewMaxHeap()

// 	heap.Insert(2)
// 	heap.Insert(5)
// 	heap.Insert(10)
// 	heap.Insert(3)

// 	fmt.Println(heap.data)

// 	maxElement, _ := heap.ExtractMax()
// 	fmt.Println("Extracted min:", maxElement)

// 	fmt.Println(heap.data)	
// }