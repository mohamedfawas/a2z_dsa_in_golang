package main

import (
	problemsarray "arraysdemo/problems"
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Left rotate array by one place
	LRotArr := problemsarray.LeftRotateByOnePlace(arr)
	fmt.Println(LRotArr)

	temp := make([]int, 3)
	fmt.Println(temp)
}
