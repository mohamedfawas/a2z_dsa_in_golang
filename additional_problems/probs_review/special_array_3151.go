package probsreview

func isArraySpecial(nums []int) bool {
	n := len(nums)
	if n == 1 {
		return true
	}

	for i := 0; i < n-1; i++ {
		currentElementIsEven := nums[i]%2 == 0
		nextElementIsEven := nums[i]%2 == 0

		if currentElementIsEven == nextElementIsEven {
			return false
		}
	}

	return true
}
