package _5_jump_game

// 最初はこっちで解いた
// Time: O(n^2)
// Space: O(n)
func canJumpDP(nums []int) bool {
	reach := make([]bool, len(nums))
	reach[0] = true

	for i, _ := range nums {
		for j := i; j >= 0; j-- {
			if nums[j] >= i-j {
				reach[i] = reach[j]
				if reach[i] {
					break
				}
			}
		}
	}
	return reach[len(nums)-1]
}

// neet code の解放見て知った。
// Time O(n)
// Space O(1)
func canJumpGreedy(nums []int) bool {
	goalIdx := len(nums) - 1

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i]+i >= goalIdx {
			goalIdx = i
		}
	}
	return goalIdx == 0
}
