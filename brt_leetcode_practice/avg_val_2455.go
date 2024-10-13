func averageValue(nums []int) int {
	var sum, count int
	for i := 0; i < len(nums); i++ {
		if nums[i]%3 == 0 && nums[i]%2 == 0 {
			sum += nums[i]
			count++
		}
	}
	if count == 0 {
		return 0
	} else {
		return (sum / count)
	}
}