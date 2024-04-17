package main

import "fmt"

func fib(num int) []int {
	arr := make([]int, num)
	for i := 0; i < num; i++ {
		if i <= 1 {
			arr[i] = i
		} else {
			arr[i] = arr[i-1] + arr[i-2]
		}
	}
	return arr
}

func main() {
	fmt.Println("Give the limit :")
	var limit int
	fmt.Scan(&limit)

	fmt.Println("fibonacci series for the given number limit is : ", fib(limit))
}
