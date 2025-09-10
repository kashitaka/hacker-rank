package _46_hand_of_straights

import "sort"

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}
	count := make(map[int]int)
	for _, v := range hand {
		count[v]++
	}

	sort.Ints(hand)
	for _, num := range hand {
		if count[num] > 0 {
			for i := num; i < num+groupSize; i++ {
				if count[i] == 0 {
					return false
				}
				count[i]--
			}
		}
	}
	return true
}
