package main

import (
	"fmt"
	bubblesort "sorting/bubble_sorting"
	selectionsorting "sorting/selection_sorting"
)

func main() {
	nums := []int{5, 6, 4, 7, 3, 44, 77, 23, 9}
	// sort
	selectionsorting.SelectionSort(nums)
	fmt.Printf("After selection sorting : %v", nums)

	nums2 := []int{5, 6, 4, 7, 3, 44, 77, 23, 9, 34}
	bubblesort.BubbleSort(nums2)
	fmt.Printf("after bubble sort is : %v", nums2)
}
