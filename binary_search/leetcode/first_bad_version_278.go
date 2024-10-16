package binsearchleetcode

func isBadVersion(num int) bool {
	// dummy function to enact the same name function shown in leetcode
	return num == num
}

func firstBadVersion(n int) int {
	low := 1
	high := n
	for low <= high {
		mid := low + (high-low)/2
		result := isBadVersion(mid)
		if result == true {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return low
}
