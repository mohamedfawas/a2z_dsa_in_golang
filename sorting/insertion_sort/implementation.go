package insertionsort

func InsertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 && arr[j-1] > arr[j] {
			arr[j-1], arr[j] = arr[j], arr[j-1]
			j--
		}
	}
}
