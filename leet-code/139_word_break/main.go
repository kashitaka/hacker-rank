package _39_word_break

// [leet, code]
//   l e e t c o d e
// 0 1 2 3 4 5 6 7 8
// T F F F T F F F T
// dp[i] = dp[i - len(word)]
// dp[8] = dp[8 - len("code")] = dp[8-4] = dp[4] = true

func wordBreak(s string, wordDict []string) bool {
	// dp[i] shows if the substring[i:] can break with wordDict
	dp := make([]bool, len(s)+1)
	dp[len(s)] = true

	for i := len(s) - 1; i >= 0; i-- {
		for _, word := range wordDict {
			wordLen := len(word)
			if i+wordLen > len(s) {
				continue
			}
			if word == s[i:i+wordLen] && dp[i+wordLen] {
				dp[i] = true
				break
			}
		}
	}
	return dp[0]
}
