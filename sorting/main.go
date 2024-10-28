package main

import (
	"fmt"
	selectionsorting "sorting/selection_sorting"
)

func main() {
	nums := []int{5, 6, 4, 7, 3, 44, 77, 23, 9}
	// sort
	selectionsorting.SelectionSort(nums)
	fmt.Printf("After selection sorting : %v", nums)
}
