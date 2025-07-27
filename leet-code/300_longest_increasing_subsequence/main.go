package _00_longest_increasing_subsequence

// neet code の回答そのもの
// むずすぎますって

func lengthOfLIS(nums []int) int {
	lis := make([]int, len(nums))
	for i, _ := range nums {
		lis[i] = 1
	}

	for i := len(nums) - 1; i >= 0; i-- {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] {
				lis[i] = max(lis[i], 1+lis[j])
			}
		}
	}
	res := 0
	for _, v := range lis {
		res = max(res, v)
	}
	return res
}
