package probsreview

/*
if len(arr)==0, return 0
if len(arr)==1, return first element
iterate from second element
	if arr[i]>arr[i-1]
		currSum += arr[i]
			if currSum>maxSum, update
	else
		currSum = arr[i]

return maxSum
*/
func maxAscendingSum(nums []int) int {
	// If the array is empty
	if len(nums) == 0 {
		return 0
	}

	// Keep track of current sum and maximum sum
	// Start from first element in the array
	currentSum := nums[0]
	maxSum := currentSum

	if len(nums) == 1 {
		return maxSum
	}

	// Iterate from second element
	for i := 1; i < len(nums); i++ {

		// If the current element is greater than previous element
		if nums[i] > nums[i-1] {
			currentSum += nums[i]

			if currentSum > maxSum {
				maxSum = currentSum
			}
		} else {
			// current element is less than prev element
			currentSum = nums[i]
		}
	}

	return maxSum
}
