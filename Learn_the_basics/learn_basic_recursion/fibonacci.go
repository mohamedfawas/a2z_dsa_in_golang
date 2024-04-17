package main

import "fmt"

func fibonacci(num int) int {
	if num <= 1 {
		//fmt.Println(num)
		return num
	}
	var last int = fibonacci(num - 1)
	var slast int = fibonacci(num - 2)
	//fmt.Println(last + slast)
	return last + slast
}

func main() {

	fmt.Println("Give the limit upto which fib numbers you want to print : ")
	var limit int
	fmt.Scan(&limit)

	//fmt.Println(fibonacci(limit))
	fibonacci(limit)
}
