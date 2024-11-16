package wordbreaktrie

func wordBreak(s string, wordDict []string) bool {
	// Convert the wordDict to a map for faster lookups
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	// Create a dp array to store whether s[0:i] can be segmented
	dp := make([]bool, len(s)+1)
	dp[0] = true // Base case: an empty string can be segmented

	// Iterate over the string
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			// Check if s[j:i] is a word and dp[j] is true
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}
