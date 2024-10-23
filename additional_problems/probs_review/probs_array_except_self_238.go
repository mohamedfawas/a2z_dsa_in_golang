package probsreview

func productExceptSelf(nums []int) []int {
	n := len(nums)
	// Initialize result array with 1s
	result := make([]int, n)
	for i := range result {
		result[i] = 1
	}

	// Step 1: Calculate prefix products
	// For each position i, compute product of all elements to its left
	// This will store in result[i] the product of all numbers to the left of nums[i]
	prefixProduct := 1
	for i := 0; i < n; i++ {
		result[i] = prefixProduct
		prefixProduct *= nums[i]
	}

	// Step 2: Calculate suffix products and combine with prefix products
	// For each position i, multiply the existing prefix product with the product
	// of all elements to its right
	suffixProduct := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= suffixProduct
		suffixProduct *= nums[i]
	}

	return result
}
