package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	array := make([]string, n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			array[i-1] = "FizzBuzz"
		} else if i%3 == 0 {
			array[i-1] = "Fizz"
		} else if i%5 == 0 {
			array[i-1] = "Buzz"
		} else {
			array[i-1] = strconv.Itoa(i)
		}
	}
	return array
}

func main() {
	// i := 2
	// ex1 := strconv.Itoa(i)
	// fmt.Println(ex1)
	arr := fizzBuzz(15)
	fmt.Println(arr)
}
