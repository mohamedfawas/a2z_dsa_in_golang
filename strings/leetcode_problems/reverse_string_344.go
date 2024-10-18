package leetcodestrings

func reverseString(s []byte) []byte {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
	return s
}

/*
func main() {
	str := "hello fawas"
	// Convert string to []byte
	reversedStr := reverseString([]byte(str))
	// Convert []byte back to string for printing
	fmt.Println(string(reversedStr))
}
*/
