package main

func numIdenticalPairs(nums []int) int {
	var count int
	numsMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		numsMap[nums[i]]++
	}
	for _, val := range numsMap {
		count += val * (val - 1) / 2
	}
	return count
}

func main() {

}
