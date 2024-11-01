package stackleetcode

func IsBalanced(s string) bool {
	stack := []rune{}

	matchingPairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		if open, isClosing := matchingPairs[char]; isClosing {
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}

			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}
