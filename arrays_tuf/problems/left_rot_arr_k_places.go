package problemsarray

// works for left rotation only

func rotate(nums []int, k int) {
	n := len(nums)

	// create temp array
	temp := make([]int, k)

	// add first k elements to temp array
	for i := 0; i < k; i++ {
		temp[i] = nums[i]
	}

	// shift the remaining elements
	for i := k; i < n; i++ {
		nums[i-k] = nums[i]
	}

	// add back temp array elements
	j := 0
	for i := n - k; i < n; i++ {
		nums[i] = temp[j]
		j++
	}
}
