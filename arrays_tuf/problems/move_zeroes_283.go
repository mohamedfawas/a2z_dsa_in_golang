package problemsarray

func moveZeroes(nums []int) {
	// Find the first occurence of zero
	j := -1
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			j = i
			break
		}
	}

	// If no occurence of zeroes
	if j == -1 {
		return
	}

	// Iterate and swap
	for i := j + 1; i < n; i++ {
		if nums[i] != 0 {
			// swap
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}
}
