package binsearchproblems

func BinarySearchRecursion(arr []int, target, low, high int) int {
	if low > high {
		return -1 // target not found
	}

	mid := low + (high-low)/2

	// Check if the target is present at mid
	if arr[mid] == target {
		return mid
	}

	// If target is smaller than mid, search in left half
	if target < arr[mid] {
		return BinarySearchRecursion(arr, target, low, mid-1)
	}

	return BinarySearchRecursion(arr, target, mid+1, high)
}
