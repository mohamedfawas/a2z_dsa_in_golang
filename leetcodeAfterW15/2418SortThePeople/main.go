package main

import "sort"

func sortPeople(names []string, heights []int) []string {
	nameHeightMap := make(map[int]string)
	for i := 0; i < len(names); i++ {
		nameHeightMap[heights[i]] = names[i]
	}
	sort.Ints(heights)
	var result []string
	for i := len(heights) - 1; i >= 0; i-- {
		result = append(result, nameHeightMap[heights[i]])
	}

	return result
}
