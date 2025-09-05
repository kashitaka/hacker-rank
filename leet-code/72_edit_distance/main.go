package _2_edit_distance

// Space: O(n)
// Time: O(nm)
// where n and m are the length of word1 and word2 respectively
func minDistance(word1 string, word2 string) int {
	dp := make([]int, len(word1)+1)
	for i, _ := range dp {
		dp[i] = len(word1[i:]) // こうした方がより説明的かも
	}

	for i := len(word2) - 1; i >= 0; i-- {
		newDP := make([]int, len(word1)+1)
		newDP[len(word1)] = len(word2[i:])
		for j := len(word1) - 1; j >= 0; j-- {
			if word1[j] == word2[i] {
				newDP[j] = dp[j+1]
			} else {
				newDP[j] = 1 + min(newDP[j+1], min(dp[j+1], dp[j]))
			}
		}
		dp = newDP
	}

	return dp[0]
}

func minDistance_naive(word1 string, word2 string) int {
	// dp[i][j] は word1[i:] と word2[j:] を一致させる最小cost
	// dp[len(word1)][len(word2)] は "" と "" を一致させるコスト = 0 (揃っている)
	//
	// dp[len(word1)][j] では "" を word2[j:]に一致させるコストであるから、
	// word2[j:]の長さ分だけコストがかかる
	// 反対に dp[i][len(word2)] では word1[i:] を "" に一致させるコストであるから
	// word1[i:] の長さ分だけコストがかかる。
	// したがって、例えば s1 = ab, s2 = cd では
	//    a b
	// c      2
	// d      1
	//    2 1 0
	// のような初期状態になる
	dp := make([][]int, len(word1)+1)
	for i, _ := range dp {
		dp[i] = make([]int, len(word2)+1)
		dp[i][len(word2)] = len(word1) - i
	}
	for j := 0; j <= len(word2); j++ {
		dp[len(word1)][j] = len(word2) - j
	}

	for i := len(word1) - 1; i >= 0; i-- {
		for j := len(word2) - 1; j >= 0; j-- {
			if word1[i] == word2[j] {
				// 先頭が一致。これは s1[i+1:] と s2[j+1:] のコストと同じである
				dp[i][j] = dp[i+1][j+1]
			} else {
				// 先頭が一致しない。何らかの操作が必要であるから +1
				// 次の3つの操作のうち、一番いいものを選ぶ
				// - replace なら一致させるため、i+1とj+1が次の sub problem
				// - remove なら i+1 と j が次の sub problem
				// - add なら i と j+1 が次のsub problem
				dp[i][j] = 1 + min(dp[i+1][j+1], min(dp[i+1][j], dp[i][j+1]))
			}
		}
	}

	return dp[0][0]
}
