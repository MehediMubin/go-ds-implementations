package main

import (
	"fmt"

	"github.com/MehediMubin/go-ds-implementations/algorithm/heapsort"
)

func main() {
	ans, _ := heapsort.HeapSort([]int{9, 5, 4, 3, 10, 1})
	fmt.Println(ans)
}