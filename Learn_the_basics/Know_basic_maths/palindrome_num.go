package main

import "fmt"

func main() {
	fmt.Println("Give the input number: ")
	var num int
	fmt.Scan(&num)

	var rev_int int
	var initial_num int = num
	for num > 0 {
		last_digit := num % 10
		rev_int = rev_int*10 + last_digit

		num /= 10
	}
	fmt.Println("The reversed integer is : ", rev_int)

	if rev_int == initial_num {
		fmt.Println("The given number is a palindrome")
	} else {
		fmt.Println("Not a palindrome")
	}
}
