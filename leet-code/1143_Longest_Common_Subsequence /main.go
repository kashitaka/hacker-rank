package _143_Longest_Common_Subsequence_


// これはむずい！！
// DPの奥深さ！
func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}

	for i := len(text1)-1; i >= 0; i-- {
		letter1 := text1[i]
		for j := len(text2)-1; j >= 0; j-- {
			letter2 := text2[j]
			if letter1 != letter2 {
				dp[i][j] = max(dp[i+1][j], dp[i][j+1])
			} else {
				dp[i][j] = dp[i+1][j+1] + 1
			}
		}
	}
	return dp[0][0]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}