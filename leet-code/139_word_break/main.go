package _39_word_break

// [leet, code]
//   l e e t c o d e
// 0 1 2 3 4 5 6 7 8
// T F F F T F F F T
// dp[i] = dp[i - len(word)]
// dp[8] = dp[8 - len("code")] = dp[8-4] = dp[4] = true

func wordBreakBackword(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[len(s)] = true

	for i := len(s) - 1; i >= 0; i-- {
		for _, word := range wordDict {
			if len(s)-i >= len(word) && s[i:len(word)+i] == word {
				dp[i] = dp[i+len(word)]
				if dp[i] {
					break
				}
			}
		}
	}
	return dp[0]
}

func wordBreakForward(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for _, word := range wordDict {
			if i-len(word) >= 0 && s[i-len(word):i] == word {
				dp[i] = dp[i-len(word)]
				if dp[i] {
					break
				}
			}
		}
	}
	return dp[len(s)]
}
