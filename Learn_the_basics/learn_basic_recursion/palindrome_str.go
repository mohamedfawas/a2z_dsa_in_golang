package main

import (
	"fmt"
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(s, "")
	s = strings.ReplaceAll(s, " ", "")
	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true

}

func main() {
	fmt.Println("test")
	fmt.Println("Give the input string : ")
	var str string = "A man, a plan, a canal: Panama"
	//fmt.Scan(&str)
	fmt.Println(isPalindrome(str))

}
