package main

import "fmt"

func largestInArray(nums []int) int {
	var max_val int = nums[0]
	for _, v := range nums {
		if v > max_val {
			max_val = v
		}
	}
	return max_val
}

func main() {
	arr := []int{5, 6, 3, 7, 3, 12, 666, 23, 55, 3}
	fmt.Println(largestInArray(arr))
}
