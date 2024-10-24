package arraysandhashing

import (
	"sort"
	"strings"
)

/*
Map to store sorted version of the word as key, value is the list of anagrams
*/
func groupAnagrams(strs []string) [][]string {
	// Create a map where the key is sorted version of the word
	// And the value is a list of words that are anagrams
	anagramMap := make(map[string][]string)

	// Iterate over each string in input slices "strs"
	for _, str := range strs {
		// Convert the string into slice of runes (so we can sort it)
		chars := strings.Split(str, "")
		// sort the slice of chars
		sort.Strings(chars)
		// Join the sorted chars back into a string
		sortedStr := strings.Join(chars, "")

		//Append the original string into the map, using the sorted string as the key
		anagramMap[sortedStr] = append(anagramMap[sortedStr], str)
	}

	// prepare a slice to hold the final group of anagrams
	var result [][]string

	// Iterate over map to collect all the groups of anagrams
	for _, group := range anagramMap {
		result = append(result, group)
	}

	return result
}
