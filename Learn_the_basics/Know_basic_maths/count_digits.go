/*
Problem Statement: Given an integer N, write a program to count the number of digits in N.
*/

package main

import "fmt"

func main() {
	fmt.Println("Give an input integer to calculate the number of digits: ")
	var input_int int
	fmt.Scan(&input_int)

	var cnt int
	for input_int > 0 {
		cnt++
		input_int /= 10
	}
	fmt.Println("Number of digits in the given interger is : ", cnt)
}
