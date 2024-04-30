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

	fmt.Println(fib(5))

}
