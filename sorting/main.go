package main

import (
	"fmt"
	bubblesort "sorting/bubble_sorting"
)

func main() {
	nums := []int{5, 6, 4, 7, 3, 44, 77, 23, 9}
	// sort
	bubblesort.BubbleSort(nums)
	fmt.Printf("After  sorting : %v", nums)
}
