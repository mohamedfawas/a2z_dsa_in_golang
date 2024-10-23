package probsreview

import "strings"

func countKeyChanges(s string) int {
	// If string is empty or has only one character, no key changes needed
	if len(s) <= 1 {
		return 0
	}

	// Initialize counter for key changes
	keyChanges := 0

	// Convert entire string to lowercase to ignore caps lock/shift
	// This simplifies comparison since 'a' and 'A' are considered same key
	s = strings.ToLower(s)

	// Loop through the string starting from index 1
	// Compare each character with the previous character
	for i := 1; i < len(s); i++ {
		// Get current and previous characters
		currentChar := s[i]
		prevChar := s[i-1]

		// If current character is different from previous character
		// it means user had to change the key
		if currentChar != prevChar {
			keyChanges++
		}
	}

	return keyChanges
}
