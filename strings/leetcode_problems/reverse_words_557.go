package leetcodestrings

import "strings"

// reverseWord takes a string and returns its reverse
// Example: "hello" -> "olleh"
func reverseWord(word string) string {
	// Convert string to rune slice to handle Unicode characters properly
	runes := []rune(word)

	// Get the length of the word
	length := len(runes)

	// Create a new rune slice to store the reversed word
	reversed := make([]rune, length)

	// Copy characters in reverse order
	for i := 0; i < length; i++ {
		reversed[i] = runes[length-1-i]
	}

	// Convert back to string and return
	return string(reversed)
}

// reverseWords takes a sentence and reverses each word while maintaining word order
func reverseWordss(s string) string {
	// Split the input string into words
	words := strings.Split(s, " ")

	// Create a slice to store reversed words
	reversedWords := make([]string, len(words))

	// Reverse each word and store in the new slice
	for i, word := range words {
		reversedWords[i] = reverseWord(word)
	}

	// Join the reversed words back together with spaces
	return strings.Join(reversedWords, " ")
}
