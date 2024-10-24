package practice

type ListNode struct {
	Value int
	Next  *ListNode
}

func FindLongestIncreasingSubArray(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}

	currentStart := 0
	currentLength := 1

	maxStart := 0
	maxLength := 1

	for i := 1; i < n; i++ {
		if arr[i] > arr[i-1] {
			currentLength++

			if currentLength > maxLength {
				maxLength = currentLength
				maxStart = currentStart
			}
		} else {
			currentStart = i
			currentLength = 1
		}
	}

	return arr[maxStart : maxStart+maxLength]
}
