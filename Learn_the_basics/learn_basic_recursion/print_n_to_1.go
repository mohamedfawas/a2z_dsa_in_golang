// Print N to 1 using recursion
package main

import "fmt"

func printNto1(num_limit int) {
	if num_limit == 0 {
		return
	}
	fmt.Println(num_limit)
	num_limit--
	printNto1(num_limit)
}

func main() {
	fmt.Println("Give the input number limit : ")
	var limit int
	fmt.Scan(&limit)
	printNto1(limit)
}
