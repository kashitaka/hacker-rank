package _851_minimum_intervals_to_include_each_query

import (
	"container/heap"
	"sort"
)

type intv struct {
	end  int
	size int
}

type MinHeap []*intv

func (m MinHeap) Len() int           { return len(m) }
func (m MinHeap) Less(i, j int) bool { return m[i].size < m[j].size }
func (m MinHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

func (m *MinHeap) Push(x any) {
	*m = append(*m, x.(*intv))
}

func (m *MinHeap) Pop() any {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}

func (m *MinHeap) Peek() *intv {
	return (*m)[0]
}

func minInterval(intervals [][]int, queries []int) []int {
	queryToRes := make(map[int]int)
	sortedQueries := make([]int, len(queries))
	copy(sortedQueries, queries)
	sort.Ints(sortedQueries)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	i := 0

	h := &MinHeap{}
	heap.Init(h)

	for _, q := range sortedQueries {
		// もう始まってるIntervalsのPush
		for i < len(intervals) && intervals[i][0] <= q {
			heap.Push(h, &intv{
				end:  intervals[i][1],
				size: intervals[i][1] - intervals[i][0] + 1,
			})
			i++
		}

		// もう終わってるIntervalsの吐き出し
		for h.Len() > 0 && h.Peek().end < q {
			heap.Pop(h)
		}

		// 解答あれば入れる
		if h.Len() > 0 {
			queryToRes[q] = h.Peek().size
		} else {
			// なければ -1
			queryToRes[q] = -1
		}

	}

	// query の order通りに答えを詰め替え
	res := make([]int, len(queries))
	for i, v := range queries {
		res[i] = queryToRes[v]
	}
	return res
}
