package main

import (
	"fmt"
	insertionsort "sorting/insertion_sort"
)

func main() {
	nums := []int{5, 6, 4, 7, 3, 44, 77, 23, 9}
	// sort
	insertionsort.InsertionSort(nums)
	fmt.Printf("After  sorting : %v", nums)
}
