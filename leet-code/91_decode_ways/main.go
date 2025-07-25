package _1_decode_ways

// これは難しすぎる。

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	dp := make([]int, len(s)+1)
	dp[0], dp[1] = 1, 1

	for i := 2; i <= len(s); i++ {
		digit1 := s[i-1]
		digit2 := s[i-2]

		if digit1 != '0' {
			// digit1 is valid encoding. so use i-1 combination
			dp[i] = dp[i-1]
		}
		if digit2 == '1' || (digit2 == '2' && '0' <= digit1 && digit1 <= '6') {
			// 10~26 is valid encoding. so use this combination
			dp[i] += dp[i-2]
		}
	}

	return dp[len(s)]
}
