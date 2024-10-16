package binsearchimplementation

func BinarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	for high >= low {
		mid := high + (high-low)/2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
