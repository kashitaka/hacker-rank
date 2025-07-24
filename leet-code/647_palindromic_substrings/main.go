package _47_palindromic_substrings

func countSubstringsCenterExpansion(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		// odd
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			res++
			l--
			r++
		}

		// even
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			res++
			l--
			r++
		}
	}

	return res
}

func countSubstringsDP(s string) int {
	dp := make([][]bool, len(s))
	for i, _ := range dp {
		dp[i] = make([]bool, len(s))
	}

	res := 0
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] && (j-i+1 <= 3 || dp[i+1][j-1]) {
				res++
				dp[i][j] = true
			}
		}
	}
	return res
}
