package main

import (
	"fmt"
	"strconv"
)

func plusOne(digits []int) []int {
	var result int
	for i := 0; i < len(digits); i++ {
		result = result*10 + digits[i]
	}
	result++
	str := strconv.Itoa(result)
	resArray := make([]int, len(str))
	for i, ch := range str {
		dig, _ := strconv.Atoi(string(ch))
		resArray[i] = dig
	}
	return resArray

}

func main() {
	nums := []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}
	fmt.Println(plusOne(nums))
	//plusOne()
}
