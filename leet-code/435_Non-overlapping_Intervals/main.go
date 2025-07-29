package _35_Non_overlapping_Intervals

import (
	"math"
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	currentEndTime := math.MinInt
	remove := 0
	for _, v := range intervals {
		if v[0] < currentEndTime {
			remove++
		}
		currentEndTime = v[1]
	}
	return remove
}

func eraseOverlapIntervalsShrinkMerge(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	currentEndTime := math.MinInt
	remove := 0
	for _, v := range intervals {
		if v[0] < currentEndTime {
			remove++
			continue
		}
		currentEndTime = v[1]
	}
	return remove
}
