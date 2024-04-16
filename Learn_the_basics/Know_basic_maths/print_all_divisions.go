package main

import "fmt"

func main() {
	fmt.Println("Give the input number limit: ")
	var num int
	fmt.Scan(&num)

	fmt.Println("Divisions of the given number is :")
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			fmt.Println(i)
		}
	}
}
