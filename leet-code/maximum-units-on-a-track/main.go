package main

import "sort"

func maximumUnits(boxTypes [][]int, truckSize int) int {
	boxCountByUnit := make(map[int]int)
	for _, v := range boxTypes {
		boxCountByUnit[v[1]] += v[0]
	}
	res := 0
	var unitList []int
	for k, _ := range boxCountByUnit {
		unitList = append(unitList, k)
	}
	sort.Slice(unitList, func(i, j int) bool {
		return unitList[i] > unitList[j]
	})

	for _, v := range unitList {
		numOfBox := boxCountByUnit[v]
		if numOfBox > truckSize {
			res += truckSize * v
			break
		}
		res += numOfBox * v
		truckSize -= numOfBox
	}
	return res
}
