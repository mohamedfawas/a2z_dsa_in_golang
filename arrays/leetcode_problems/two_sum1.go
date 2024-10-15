package leetcodeproblems

func twoSum(nums []int, target int) []int {
	numRecord := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if _, found := numRecord[complement]; found {
			return []int{numRecord[complement], i}
		}
		numRecord[num] = i
	}

	return []int{}
}
