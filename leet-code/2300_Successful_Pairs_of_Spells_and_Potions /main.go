package _300_Successful_Pairs_of_Spells_and_Potions_

import "sort"

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Slice(potions, func(i, j int) bool {
		return potions[i] < potions[j]
	})

	resultCache := make(map[int]int)
	res := make([]int, len(spells))
	for i, spell := range spells {
		result, ok := resultCache[spell]
		if ok {
			res[i] = result
			continue
		}

		if spell * potions[0] >= int(success) {
			res[i] = len(potions)
			resultCache[spell]= res[i]
			continue
		}


		// binary search
		left, right := 0, len(potions)
		for left + 1 < right {
			middle := (left + right)/2
			if spell * potions[middle] < int(success) {
				left = middle
			} else {
				right = middle
			}
		}

		res[i] = len(potions) - right
		resultCache[spell]= res[i]
	}
	return res
}

