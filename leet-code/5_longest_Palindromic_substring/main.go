package __longest_Palindromic_substring

func longestPalindromeDP(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	for i, _ := range dp {
		dp[i] = make([]bool, n)
	}

	res := ""
	maxRes := 0
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n-1; j++ {
			if s[i] == s[j] && (j-i <= 2 || dp[i+1][j-1]) {
				dp[i][j] = true
				if j-i+1 > maxRes {
					res = s[i : j+1]
					maxRes = j - i + 1
				}
			}
		}
	}
	return res
}

func longestPalindromeCenterExpansion(s string) string {
	res := ""
	lenMax := 0
	for i := 0; i < len(s); i++ {
		// odd
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lenMax {
				lenMax = r - l + 1
				res = s[l : r+1]
			}
			l, r = l-1, r+1
		}

		// even
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > lenMax {
				lenMax = r - l + 1
				res = s[l : r+1]
			}
			l, r = l-1, r+1
		}
	}
	return res
}
