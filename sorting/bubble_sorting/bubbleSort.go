package bubblesort

func BubbleSort(arr []int) {
	n := len(arr)
	for i := n - 1; i <= 1; i++ {
		for j := 0; j <= i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
