package leetcodestrings

import "strings"

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	strArray := strings.Split(s, " ")
	return len(strArray[len(strArray)-1])
}
