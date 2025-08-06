package heapsort

import (
	"github.com/MehediMubin/go-ds-implementations/data-structure/heap"
)

func HeapSort(arr []int) ([]int, error) {
	h := heap.NewMinHeap()

	for _, val := range arr {
		h.Insert(val)
	}

	var sortedArray []int
	for !h.IsEmpty() {
		min, _ := h.ExtractMin()
		sortedArray = append(sortedArray, min)	
	}

	return sortedArray, nil
}

// func main() {
// 	arr, _ := HeapSort([]int{5, 10, 3, 4, 2, 1})
// 	fmt.Println(arr)
// }