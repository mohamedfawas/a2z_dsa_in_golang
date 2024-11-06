package stackleetcode

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	// Dictionary to map each element in nums2 to its next greater element
	nextGreater := make(map[int]int)
	// Stack to keep track of elements in nums2 for finding the next greater elements
	stack := []int{}

	// Traverse nums2 in reverse to populate nextGreater map
	for i := len(nums2) - 1; i >= 0; i-- {
		current := nums2[i]
		// Maintain the stack such that it only contains elements greater than the current element
		for len(stack) > 0 && stack[len(stack)-1] <= current {
			stack = stack[:len(stack)-1] // Pop the top of the stack
		}

		// If stack is not empty, the top element is the next greater element
		if len(stack) > 0 {
			nextGreater[current] = stack[len(stack)-1]
		} else {
			// If stack is empty, there is no greater element to the right
			nextGreater[current] = -1
		}

		// Push the current element onto the stack
		stack = append(stack, current)
	}

	// Build the result array for nums1 based on the nextGreater map
	result := make([]int, len(nums1))
	for i, num := range nums1 {
		result[i] = nextGreater[num]
	}

	return result
}
