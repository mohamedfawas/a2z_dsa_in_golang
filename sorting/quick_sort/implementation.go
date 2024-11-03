package quicksort

// approach shown in jenny's lecture

func partition(arr []int, lb, ub int) int {
	pivot := arr[lb]
	start := lb
	end := ub
	for start < end {
		for start < ub && arr[start] <= pivot {
			start++
		}
		for arr[end] > pivot {
			end--
		}
		//swap start and end
		if start < end {
			arr[start], arr[end] = arr[end], arr[start]
		}
	}

	//swap pivot and end
	arr[end], arr[lb] = arr[lb], arr[end]
	return end
}

func QuickSort(arr []int, lb, ub int) {
	if lb < ub {
		loc := partition(arr, lb, ub)
		QuickSort(arr, lb, loc-1)
		QuickSort(arr, loc+1, ub)
	}
}
