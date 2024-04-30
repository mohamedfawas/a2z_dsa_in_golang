package main

import (
	"fmt"
	"math"
)

func secondLargest(nums []int) int {
	var largest, s_largest int
	largest = math.MinInt64
	for _, v := range nums {
		if v > largest {
			s_largest = largest
			largest = v
		} else if v > s_largest && v != largest {
			s_largest = v
		}
	}
	return s_largest
}

func main() {
	arr := []int{5, 6, 3, 7, 3, 12, 666, 23, 55, 3}
	fmt.Println(secondLargest(arr))
	//fmt.Println(largestInArray(arr))
}
