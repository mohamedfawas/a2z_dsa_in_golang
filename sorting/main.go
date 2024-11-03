package main

import (
	"fmt"
	quicksort "sorting/quick_sort"
)

func main() {
	nums := []int{5, 6, 4, 7, 3, 44, 77, 23, 9}
	// sort
	quicksort.QuickSort(nums, 0, len(nums)-1)
	fmt.Printf("After  sorting : %v", nums)
}
