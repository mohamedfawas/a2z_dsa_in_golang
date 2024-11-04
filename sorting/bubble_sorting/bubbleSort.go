package bubblesort

func BubbleSort(arr []int) {
	/*
	 */
	n := len(arr)
	for i := 0; i < n-1; i++ {
		flag := 0
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = 1
			}
		}
		if flag == 0 {
			break
		}
	}
}
