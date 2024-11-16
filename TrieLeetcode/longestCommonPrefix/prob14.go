package longestcommonprefixtrie

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for _, str := range strs[1:] {

		for len(prefix) > 0 && !startsWith(str, prefix) {
			prefix = prefix[:len(prefix)-1]
		}

		if prefix == "" {
			return ""
		}
	}

	return prefix
}

func startsWith(str, prefix string) bool {
	if len(prefix) > len(str) {
		return false
	}
	return str[:len(prefix)] == prefix
}
