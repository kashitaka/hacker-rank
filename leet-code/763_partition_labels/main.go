package _63_partition_labels

func partitionLabels(s string) []int {
	lastAppear := make(map[rune]int)
	for i, v := range s {
		lastAppear[v] = i
	}

	start, end := 0, -1
	res := []int{}

	for i, v := range s {
		end = max(end, lastAppear[v])
		if i == end {
			res = append(res, end-start+1)
			start = i + 1
		}
	}
	return res
}
