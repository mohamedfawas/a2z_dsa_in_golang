package implementation

func LinearSearch(target int, arr []int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i
		}
	}

	return -1
}
