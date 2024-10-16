package binsearchleetcode

func guess(n int) int {
	// this is a dummy function to enact the function given in leetcode problem
	return 1
}

func guessNumber(n int) int {
	/*
		low=1
		high=n
		for low<=high{
			find mid
			call function(mid int)
			compare output
			if -1:
				high:= mid-1
			1:
				low := mid+1
			0:
			 mid
		}
	*/
	low := 1
	high := n
	for low <= high {
		mid := low + (high-low)/2
		output := guess(mid)
		if output == 0 {
			return mid
		} else if output == -1 {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
