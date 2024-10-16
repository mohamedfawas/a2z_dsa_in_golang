package linsearchleetcode

func findNumbers(nums []int) int {
	var count int
	for _, num := range nums {
		var num_digit int
		for num > 0 {
			num /= 10
			num_digit++
		}
		if num_digit%2 == 0 {
			count++
		}
	}
	return count
}
