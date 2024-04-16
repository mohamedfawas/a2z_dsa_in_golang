// Sum of first N numbers
package main

import "fmt"

func countFunction(number int, sum int) {
	if number < 1 {
		fmt.Println("sum of the numbers is : ", sum)
		return
	}
	sum += number
	number--
	countFunction(number, sum)
}

func main() {
	fmt.Println("give the input number : ")
	var num int
	fmt.Scan(&num)

	var sum int
	countFunction(num, sum)

}
