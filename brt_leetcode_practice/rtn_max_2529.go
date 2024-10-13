package main

func maximumCount(nums []int) int {
	var countP, countN int
	for _, val := range nums {
		if val > 0 {
			countP++
		} else if val < 0 {
			countN++
		}
	}
	return max(countP, countN)
}

func main() {

}
