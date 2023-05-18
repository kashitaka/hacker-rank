package _92_top_k_grequent_words

import "sort"

func topKFrequent(words []string, k int) []string {
	freq := make(map[string]int)
	var uniq []string

	for _, v := range words {
		_, ok := freq[v]
		if !ok {
			uniq = append(uniq, v)
		}
		freq[v]++
	}

	sort.Slice(uniq, func(i, j int) bool {
		if freq[uniq[i]] == freq[uniq[j]] {
			return uniq[i] < uniq[j] // 辞書昇順
		}
		return freq[uniq[i]] > freq[uniq[j]] // 数字降順
	})

	return uniq[:k]
}