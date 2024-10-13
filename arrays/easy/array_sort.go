package main

import "fmt"

func checkSorted(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] {
			return false
		}
	}
	return true
}

func main() {
	arr := []int{1, 2, 1, 4, 6}
	fmt.Println(checkSorted(arr))
}
