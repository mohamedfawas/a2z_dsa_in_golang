package main

func numIdenticalPairs(nums []int) int {
	freqCount := make(map[int]int)
	for _, num := range nums {
		freqCount[num]++
	}

	pairCount := 0
	for _, count := range freqCount {
		pairCount += count * (count - 1) / 2
	}

	return pairCount
}
