package queueleetcode

func firstUniqChar(s string) int {
	// Step 1: Create a map to count the occurrences of each character.
	// This will help us quickly check how many times each character appears in the string.
	charCount := make(map[rune]int)

	// Step 2: Use a queue to store the characters in the order they appear in the string.
	// This helps us track the order of non-repeating characters as we process them.
	type charIndex struct {
		char  rune
		index int
	}
	queue := []charIndex{}

	// Step 3: Loop through the string and populate both the count map and queue.
	for i, char := range s {
		// Increment the character count in the map
		charCount[char]++

		// Add the character and its index to the queue
		// This keeps the character order intact.
		queue = append(queue, charIndex{char, i})
	}

	// Step 4: Process the queue to find the first non-repeating character.
	// We keep removing characters from the front of the queue until we find one that has a count of 1.
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:] // Remove the front element from the queue

		// Check if this character is non-repeating
		if charCount[front.char] == 1 {
			return front.index // Return the index of the first non-repeating character
		}
	}

	// Step 5: If we exit the loop, it means there are no non-repeating characters in the string.
	return -1
}
