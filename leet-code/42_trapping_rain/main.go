package _2_trapping_rain

func trap(height []int) int {
	maxLeft := make([]int, len(height))
	maxRight := make([]int, len(height))
	maxLeft[0], maxRight[len(maxRight)-1] = 0, 0

	for i := 1; i < len(height); i++ {
		maxLeft[i] = max(maxLeft[i-1], height[i-1])
	}
	for i := len(height) - 2; i >= 0; i-- {
		maxRight[i] = max(maxRight[i+1], height[i+1])
	}
	res := 0
	for i := 0; i < len(height); i++ {
		res += max(min(maxLeft[i], maxRight[i])-height[i], 0)
	}
	return res
}
