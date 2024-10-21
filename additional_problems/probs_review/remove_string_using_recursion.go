package probsreview

/*
remove character from string using recursion
*/

func RemoveChar(s string, target rune) string {
	// Base case : if the string is empty return an empty string
	if len(s) == 0 {
		return ""
	}

	// Check if the first character of the string is target character
	if rune(s[0]) == target {
		// If it matches
		return RemoveChar(s[1:], target)
	}

	// If the first character does not match, keep it and call the function with the rest of the string.
	return string(s[0]) + RemoveChar(s[1:], target)
}
