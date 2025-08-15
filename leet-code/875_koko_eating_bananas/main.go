package _75_koko_eating_bananas

import "math"

func minEatingSpeed(piles []int, h int) int {
	maxK := 0
	for _, v := range piles {
		maxK = max(maxK, v)
	}

	l, r := 0, maxK
	for l+1 < r {
		mid := (l + r) / 2
		t := 0.0
		for _, v := range piles {
			t += math.Ceil(float64(v) / float64(mid))
		}
		if t <= float64(h) {
			// koko is fast and may take more time. so search left
			// 仮にhと同じでももっと遅くできるかもしれない。
			r = mid
		} else {
			// koko is too slow. so search right
			l = mid
		}
	}
	return r
}
