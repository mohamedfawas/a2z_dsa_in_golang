package main

func largestOddNumber(num string) string {
	// Start at the last character of the string and work backward.
	// This is because the largest odd number must end at the last odd digit we find.
	for i := len(num) - 1; i >= 0; i-- {
		// Convert the current character to its numeric value.
		// To do this, we subtract '0' (ASCII value of 48) from the character.
		// Example: '5' - '0' = 5, which is the number 5.
		if (num[i]-'0')%2 != 0 {
			// Check if this number is odd.
			// A number is odd if it leaves a remainder of 1 when divided by 2.
			// Example: 5%2 = 1, so it's odd.

			// If we find an odd number, we return everything from the start
			// of the string up to this character (inclusive).
			// This is because any substring ending at this position is larger
			// than substrings ending earlier.
			return num[:i+1]
		}
	}
	// If we finish the loop and don't find any odd numbers, return an empty string.
	// This means the input string only contains even digits.
	return ""
}
