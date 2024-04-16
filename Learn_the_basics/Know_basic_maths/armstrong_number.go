package main

import "fmt"

func main() {
	fmt.Println("Give the input number: ")
	var num int
	fmt.Scan(&num)

	var final_sum int
	var initial_num int = num
	for num > 0 {
		last_digit := num % 10
		final_sum = final_sum + (last_digit * last_digit * last_digit)

		num /= 10
	}
	fmt.Println("sum of cubes of individual digits is : ", final_sum)

	if initial_num == final_sum {
		fmt.Println("It's an armstrong number")
	} else {
		fmt.Println("Not an armstrong number")
	}

}
