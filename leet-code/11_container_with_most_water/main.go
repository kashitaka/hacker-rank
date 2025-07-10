package main

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	res := 0
	for l < r {
		left := height[l]
		right := height[r]
		minHeight := left
		dist := r - l
		if right < left {
			minHeight = right
			r--
		} else {
			l++
		}
		area := dist * minHeight
		if area > res {
			res = area
		}
	}
	return res
}
