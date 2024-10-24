package main

import (
	"fmt"
)

// func reverseWords(s string) string {
// 	if len(s) <= 1 {
// 		return s
// 	}

// 	return reverseWords(s[1:]) + string(s[0])
// }

func main() {
	word := "hello fawas"
	fmt.Println(reverseWords(word))
}
