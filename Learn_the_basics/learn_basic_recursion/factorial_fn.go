package main

import "fmt"

func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}

func main() {
	fmt.Println("give the input number :")
	var num int
	fmt.Scan(&num)
	fact := factorial(num)
	fmt.Println("factorial of the given number is : ", fact)

}
