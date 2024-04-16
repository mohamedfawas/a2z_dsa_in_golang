// greatest common divisor
package main

import "fmt"

func main() {

	var num1, num2 int
	fmt.Println("give the input number num1 :")
	fmt.Scan(&num1)
	fmt.Println("give the input number num2 :")
	fmt.Scan(&num2)
	fmt.Println("Minimum of the given two numbers is : ", min(num1, num2))

	var gcd int
	for i := min(num1, num2); i >= 1; i-- {
		if num1%i == 0 && num2%i == 0 {
			gcd = i
			break
		}
	}
	fmt.Println("gcd of the given two numbers is : ", gcd)
}
