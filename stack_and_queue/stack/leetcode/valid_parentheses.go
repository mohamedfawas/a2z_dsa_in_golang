package stackleetcode

func isValid(s string) bool {
	// Initialize a stack to keep track of opening brackets
	stack := make([]rune, 0)

	// Create a map to store matching pairs of brackets
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		// if the character is an opening bracket
		if char == '(' || char == '{' || char == '[' {
			//  push into stack
			stack = append(stack, char)
		} else if char == ')' || char == '}' || char == ']' {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
