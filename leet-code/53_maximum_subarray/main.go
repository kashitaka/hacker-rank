package _3_maximum_subarray

func maxSubArrayKadane(nums []int) int {
	maxSum := nums[0]
	curSum := nums[0]

	for i := 1; i < len(nums); i++ {
		// 引き継ぐか、新たに始めるか
		curSum = max(curSum+nums[i], nums[i])
		// 記録更新したか
		maxSum = max(maxSum, curSum)
	}
	return maxSum
}
