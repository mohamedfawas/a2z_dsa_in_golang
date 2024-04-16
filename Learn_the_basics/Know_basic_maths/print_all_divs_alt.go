package main

import "fmt"

func main() {
	fmt.Println("Give the input number limit :")
	var num int
	fmt.Scan(&num)

	fmt.Println("The divisions of the given number is :")
	for i := 1; i*i <= num; i++ {
		if num%i == 0 {
			fmt.Println(i)
			if num/i != i {
				fmt.Println(num / i)
			}
		}
	}
}
