package recursionleetcode

func isPowerOfThree(n int) bool {
	if n == 1 {
		return true
	}
	if n%3 != 0 || n <= 0 {
		return false
	}
	return isPowerOfThree(n / 3)
}
