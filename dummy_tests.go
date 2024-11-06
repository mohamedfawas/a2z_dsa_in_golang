package main

import "fmt"

func ReverseStringUsingStack(input string) string {
	stack := []rune{}

	for _, char := range input {
		stack = append(stack, char)
	}

	reversed := ""
	for len(stack) > 0 {
		n := len(stack)
		top := stack[n-1]
		reversed += string(top)
		stack = stack[:n-1]
	}

	return reversed
}

func main() {
	str := "hello"
	fmt.Println(ReverseStringUsingStack(str))
}
