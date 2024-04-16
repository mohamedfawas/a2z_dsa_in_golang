// Print linearly from 1 to N using a recursion function
package main

import "fmt"

var count int = 1

func printLinear(limit_number int) {
	if count > limit_number {
		return
	}
	fmt.Println(count)
	count++
	printLinear(limit_number)
}

func main() {
	fmt.Println("Give the number limit N : ")
	var limit int
	fmt.Scan(&limit)
	printLinear(limit)
}
