package stackleetcode

func ReverseStringUsingStack(input string) string {
	stack := []rune{}

	for _, char := range input {
		stack = append(stack, char)
	}

	reversed := ""
	for len(stack) > 0 {
		n := len(stack) - 1
		reversed += string(stack[n])
		stack = stack[:n]
	}

	return reversed
}
