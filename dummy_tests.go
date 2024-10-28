package main

// func reverseWords(s string) string {
// 	if len(s) <= 1 {
// 		return s
// 	}

// 	return reverseWords(s[1:]) + string(s[0])
// }

func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	for i := range result {
		result[i] = 1
	}

	prefixProduct := 1
	for i := 0; i < n; i++ {
		result[i] = prefixProduct
		prefixProduct *= nums[i]
	}

	suffixProduct := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= suffixProduct
		suffixProduct *= nums[i]
	}

	return result
}

func main() {

}
