package _28_longest_consecutive_sequence

func longestConsecutive(nums []int) int {
	hash := make(map[int]bool)
	for _, v := range nums {
		hash[v] = true
	}
	res := 0
	for k, _ := range hash {
		if !hash[k-1] {
			// k is starting number
			i := 0
			for hash[k+i] {
				i++
				res = max(i, res)
			}
		}
	}
	return res
}
