package main

func countEven(num int) int {
	count := 0
	for i := 1; i < num; i++ {
		if isEven(i) {
			count++
		}
	}

	return count
}

func isEven(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}
