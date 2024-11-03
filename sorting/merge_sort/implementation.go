package mergesort

func MergeSort(arr []int) []int {
	n := len(arr)
	// Base case : if the array has 1 or 0 elements, it is already sorted
	if n <= 1 {
		return arr
	}

	// step 1 : find the middle index
	mid := n / 2

	// step 2 : recursively split and sort the left half
	left := MergeSort(arr[:mid])

	// step 3; recursively spli and sort the right half
	right := MergeSort(arr[mid:])

	// step 4 : merge the sorted halfs and return the result
	return Merge(left, right)
}

func Merge(left, right []int) []int {
	// create a slice to hold the merged result
	result := []int{}

	// Indices for iterating over the left and right slices
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	for i < len(left) {
		result = append(result, left[i])
		i++
	}
	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}
