package main

import (
	"fmt"
	//"strings"
)

func isPalindrome(str string, start int, end int) bool {
	if start < end {
		if str[start] != str[end] {
			return false
		}
		isPalindrome(str, start+1, end-1)
	}
	return true
}

func main() {
	fmt.Println("Give an input string :")
	var test_string string
	fmt.Scan(&test_string)
	start := 0
	end := len(test_string) - 1
	result := isPalindrome(test_string, start, end)
	fmt.Println(result)
}
