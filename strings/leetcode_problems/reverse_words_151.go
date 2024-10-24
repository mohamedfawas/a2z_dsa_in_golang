package leetcodestrings

import "strings"

func reverseWords(s string) string {
	// Step 1: Trim any leading or trailing spaces from the input string
	s = strings.TrimSpace(s)

	// Step 2: Split the string into words
	// strings.Fields handles multiple spaces between words and removes them
	words := strings.Fields(s)

	// Step 3: Initialize two pointers for reversing the array of words
	left := 0
	right := len(words) - 1

	// Step 4: Reverse the array of words
	for left < right {
		// Swap words at left and right pointers
		words[left], words[right] = words[right], words[left]
		// Move pointers towards the center
		left++
		right--
	}

	// Step 5: Join the words back together with a single space between them
	// strings.Join will add exactly one space between each word
	return strings.Join(words, " ")
}
