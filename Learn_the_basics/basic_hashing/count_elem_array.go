package main

import "fmt"

func preStore(nums_array []int) map[int]int {
	freq_map := make(map[int]int)
	for i := 0; i < len(nums_array)+1; i++ {
		freq_map[i]++
	}
	return freq_map
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	fmt.Println("Frequency count : ", preStore(arr))
}
