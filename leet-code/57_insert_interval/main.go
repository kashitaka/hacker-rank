package _7_insert_interval

// neet code の説明見て解いた。
// すごいシンプルだが奥深い。

func insert(intervals [][]int, newInterval []int) [][]int {
	res := [][]int{}

	for i, interval := range intervals {
		if newInterval[1] < interval[0] {
			res = append(res, newInterval)
			res = append(res, intervals[i:]...)
			return res
		} else if interval[1] < newInterval[0] {
			res = append(res, interval)
		} else {
			newInterval[0] = min(newInterval[0], interval[0])
			newInterval[1] = max(newInterval[1], interval[1])
		}
	}
	res = append(res, newInterval)
	return res
}
