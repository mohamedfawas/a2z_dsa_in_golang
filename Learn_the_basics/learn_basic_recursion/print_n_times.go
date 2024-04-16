package main

import "fmt"

var counter int

func printFunction(number int) {
	if number == counter {
		return
	}
	fmt.Println("Hai fawas")
	counter++
	printFunction(number)
}

func main() {
	fmt.Println("give the number limit :")
	var num int
	fmt.Scan(&num)
	printFunction(num)
}
