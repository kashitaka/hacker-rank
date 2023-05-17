package _46_min_cost_climing_stairs

func minCostClimbingStairs(cost []int) int {
	minStepMap := make(map[int]int, len(cost))
	minStepMap[0] = cost[0]
	minStepMap[1] = cost[1]
	for i := 2; i < len(cost); i++ {
		step2 := minStepMap[i-2] + cost[i]
		step1 := minStepMap[i-1] + cost[i]
		if step2 > step1 {
			minStepMap[i] = step1
		} else {
			minStepMap[i] = step2
		}
	}
	if minStepMap[len(cost)-1] > minStepMap[len(cost)-2] {
		return minStepMap[len(cost)-2]
	}
	return minStepMap[len(cost)-1]
}