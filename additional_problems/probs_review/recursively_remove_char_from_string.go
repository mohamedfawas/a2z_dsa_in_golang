package probsreview

// Recursive function to remove a specific character from a string
func removeChar(s string, ch rune) string {
	// Base case: if the string is empty, return an empty string
	if len(s) == 0 {
		return ""
	}

	// Recursively process the rest of the string
	restOfString := removeChar(s[1:], ch)

	// If the first character is the one to remove, return only the rest of the string
	if rune(s[0]) == ch {
		return restOfString
	}

	// Otherwise, return the first character concatenated with the rest of the string
	return string(s[0]) + restOfString
}
