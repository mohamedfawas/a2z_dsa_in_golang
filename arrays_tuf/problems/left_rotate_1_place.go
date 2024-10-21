package problemsarray

func LeftRotateByOnePlace(arr []int) []int {
	temp := arr[0]
	n := len(arr)

	for i := 1; i < n; i++ {
		arr[i-1] = arr[i]
	}

	arr[n-1] = temp

	return arr
}
