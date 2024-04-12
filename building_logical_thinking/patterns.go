package main

import "fmt"

var size int = 5

func main() {
	square_pattern()
	increasing_pattern()
	increaing_by_double()
}

func square_pattern() {
	fmt.Println("Square pattern is printed below: ")
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Printf("* ")
		}
		fmt.Println("")
	}
}

func increasing_pattern() {
	fmt.Println("Increasing pattern is printed below: ")
	for i := 1; i <= size; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("*")
		}
		fmt.Println("")
	}

}

func increaing_by_double() {
	fmt.Println("increasing by double pattern is printed below: ")
	for i := 1; i <= size; i++ {
		for j := 1; j <= 2*i; j++ {
			fmt.Printf("* ")
		}
		fmt.Println("")
	}
}
