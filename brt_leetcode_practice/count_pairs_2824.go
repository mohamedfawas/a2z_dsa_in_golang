package main

import (
	"fmt"
	"sort"
)

func countPairs(nums []int, target int) (ans int) {
	sort.Ints(nums)
	for j, v := range nums {
		i := sort.SearchInts(nums[:j], target-v)
		ans += i
	}
	fmt.Println(nums)
	return
}

func main() {
	nums := []int{-1, 1, 2, 3, 1}
	target := 2
	countPairs(nums, target)

}
