package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Println("Give the input number num1 :")
	fmt.Scan(&num1)
	fmt.Println("Give the input number num2 :")
	fmt.Scan(&num2)

	for num1 > 0 && num2 > 0 {
		if num1 > num2 {
			num1 = num1 % num2
		} else {
			num2 = num2 % num1
		}
	}

	var gcd int
	if num1 == 0 {
		gcd = num2
	} else {
		gcd = num1
	}

	fmt.Println("the gcd of the given numbers is : ", gcd)
}
