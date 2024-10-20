package arraysandhashing

func longestConsecutive(nums []int) int {
	// If the array is empty, return 0
	if len(nums) == 0 {
		return 0
	}

	// Create a map to store all numbers from input array
	// Using map for O(1) lookup time
	numMap := make(map[int]bool)

	// Add all numbers to the map
	// The boolean value doesn't matter, we just need the keys
	for _, num := range nums {
		numMap[num] = true
	}

	longestStreak := 0

	// Iterate through each number in the original array
	for _, num := range nums {
		// Check if this number could be the start of a sequence
		// If num-1 exists in map, then this number is not the start of a sequence
		if !numMap[num-1] {
			// This number could be the start of a sequence
			// Let's count how long this sequence could be
			currentNum := num
			currentStreak := 1

			// Keep checking for consecutive numbers
			// While the next number exists in our map
			for numMap[currentNum+1] {
				currentNum++
				currentStreak++
			}

			// Update longest streak if current streak is longer
			if currentStreak > longestStreak {
				longestStreak = currentStreak
			}
		}
	}

	return longestStreak
}
