package _0_climbing_stairs

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	cache := make(map[int]int)
	cache[1] = 1
	cache[2] = 2
	for i:=3; i<=n; i++ {
		// to reach step i, there are 2 ways
		// - 2 steps up from i-2 stair
		// - 1 step up from i-1 stair
		cache[i] = cache[i-2] + cache[i-1]
	}
	return cache[n]
}