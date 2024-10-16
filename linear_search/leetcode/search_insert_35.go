package linsearchleetcode

func searchInsert(nums []int, target int) int {
	for idx, num := range nums {
		if num == target || num > target {
			return idx
		}
	}
	return len(nums)
}
