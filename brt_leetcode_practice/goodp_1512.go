package main

/*
Count how many times each number appears.
If a number appears n times, then n * (n – 1) // 2 good pairs can be made with this number.
*/

func numIdenticalPairs(nums []int) int {
	var count int // to store number of good pairs
	numsMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		numsMap[nums[i]]++
	}
	for _, val := range numsMap {
		count += val * (val - 1) / 2
	}
	return count

}

func main() {

}
