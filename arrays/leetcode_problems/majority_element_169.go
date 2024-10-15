package leetcodeproblems

func majorityElement(nums []int) int {
	count := 0
	var element int
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			element = nums[i]
			count++
		} else if nums[i] == element {
			count++
		} else {
			count--
		}
	}
	return element
}
