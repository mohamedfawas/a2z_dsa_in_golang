package probsreview

import "sort"

func topKFrequent(nums []int, k int) []int {
	// Step 1: Create a map to store the frequency of each number
	// The key is the number and the value is how many times it appears
	freqMap := make(map[int]int)

	// Count the frequency of each number in the array
	for _, num := range nums {
		freqMap[num]++
	}

	// Step 2: Create a slice to store unique numbers
	// We'll sort this slice based on frequencies
	uniqueNums := make([]int, 0, len(freqMap))
	for num := range freqMap {
		uniqueNums = append(uniqueNums, num)
	}

	// Step 3: Sort the unique numbers based on their frequencies
	// We use sort.Slice with a custom less function
	// This will sort in descending order (higher frequencies first)
	sort.Slice(uniqueNums, func(i, j int) bool {
		// If frequencies are equal, compare the numbers to ensure stable sort
		if freqMap[uniqueNums[i]] == freqMap[uniqueNums[j]] {
			return uniqueNums[i] < uniqueNums[j]
		}

		return freqMap[uniqueNums[i]] > freqMap[uniqueNums[j]]
	})

	// Step 4: Return the first k elements
	return uniqueNums[:k]
}
