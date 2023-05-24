package _16_Combination_Sum_3

func combinationSum3(k int, n int) [][]int {
	var ans [][]int
	recursive(k, n, []int{}, 0, &ans, 1)
	return ans
}

func recursive(k int, n int, currElem []int, curSum int, ans *[][]int, index int) {
	if curSum > n {
		return
	}
	if len(currElem) == k && curSum == n {
		x := make([]int, k)
		copy(x, currElem)
		*ans = append(*ans, x)
	}

	for i:=index; i <=9; i++ {
		currElem = append(currElem, i)
		recursive(k, n, currElem, curSum+i, ans, i+1)
		currElem = currElem[:len(currElem)-1]
	}
}
