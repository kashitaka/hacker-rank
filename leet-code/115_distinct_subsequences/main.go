package _15_distinct_subsequences

// 素朴なやり方
func numDistinct2(s string, t string) int {
	// dp[i][j] は s[i:] と t[j:] の組み合わせのサブプロブレムの答え
	// dp[i][len(j)]とは任意の文字列を使って、空白文字を作る組み合わせ
	// 何選ばないがあるだけなので 1 を埋める
	// dp[0][0]を取り出すことで、求めている答えを得ることができる
	dp := make([][]int, len(s)+1)
	for i, _ := range dp {
		dp[i] = make([]int, len(t)+1)
		dp[i][len(t)] = 1
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := len(t) - 1; j >= 0; j-- {
			// i 番目の文字を使わない（えない）ケースでは
			// 常に i+1 番目を使った subproblem の答えを使う
			// s[i:] = ab, t[j:] = b としたときに,
			// 先頭aを使わない(えない)ので、s[i_1:] = b と b の問題を解く
			dp[i][j] = dp[i+1][j]

			// これが一致する場合は先頭文字を使える
			// その場合は s[i+1:] と t[j+1:]の subproblem を引き継ぐ
			if s[i] == t[j] {
				dp[i][j] += dp[i+1][j+1]
			}
		}
	}
	return dp[0][0]
}

// memory 最適化
func numDistinct(s string, t string) int {
	dp := make([]int, len(s)+1)
	for i, _ := range dp {
		dp[i] = 1
	}

	for i := len(t) - 1; i >= 0; i-- {
		newDP := make([]int, len(s)+1)
		for j := len(s) - 1; j >= 0; j-- {
			newDP[j] = newDP[j+1]
			if s[j] == t[i] {
				newDP[j] += dp[j+1]
			}
		}
		dp = newDP
	}
	return dp[0]
}
