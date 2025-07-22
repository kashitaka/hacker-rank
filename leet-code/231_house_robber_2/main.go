package _31_house_robber_2

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Time: O(N)
// Space: O(N)
// where N is length of nums
func firstSolution(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	robFirst, noRobFirst := make([]int, len(nums)), make([]int, len(nums))
	robFirst[0], robFirst[1] = nums[0], nums[0]
	noRobFirst[0], noRobFirst[1] = 0, nums[1]

	for i := 2; i < len(nums); i++ {
		robFirst[i] = max(robFirst[i-1], robFirst[i-2]+nums[i])
		noRobFirst[i] = max(noRobFirst[i-1], noRobFirst[i-2]+nums[i])
	}
	robFirst[len(nums)-1] = robFirst[len(nums)-2]
	return max(robFirst[len(nums)-1], noRobFirst[len(nums)-1])
}

func robNeetcodeSolution(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return max(rob1(nums[1:]), rob1(nums[:len(nums)-1]))
}

func rob1(nums []int) int {
	rob1, rob2 := 0, 0
	for _, v := range nums {
		newRob := max(rob2, v+rob1)
		rob1 = rob2
		rob2 = newRob
	}
	return rob2
}
