package problemsarray

func Rotate(nums []int, k int) {
	n := len(nums)

	// If k is 0 or array is empty, no rotation needed
	if k == 0 || n == 0 {
		return
	}

	// temp arr
	temp := make([]int, k)

	// copy last k elements to temp array
	j := 0
	for i := n - k; i < n; i++ {
		temp[j] = nums[i]
		j++
	}

	// shift the remaining elements to the right
	for i := n - k - 1; i >= 0; i-- {
		nums[i+k] = nums[i]
	}

	// copy back elements from temp to the beginning
	for i := 0; i < k; i++ {
		nums[i] = temp[i]
	}
}
