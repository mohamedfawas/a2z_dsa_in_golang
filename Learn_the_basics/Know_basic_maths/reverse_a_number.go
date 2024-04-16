package main

import "fmt"

func main() {
	var input int
	fmt.Println("Give the input integer: ")
	fmt.Scan(&input)

	var rev_int int
	for input > 0 {
		last_digit := input % 10
		rev_int = rev_int*10 + last_digit

		input /= 10
	}
	fmt.Println("The reversed integer is : ", rev_int)
}
