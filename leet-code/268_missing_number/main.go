package _68_missing_number

func missingNumberBit(nums []int) int {
	res := 0
	for _, v := range nums {
		res = res ^ v
	}
	for i := 0; i <= len(nums); i++ {
		res = res ^ i
	}
	return res
}

func missingNumberHash(nums []int) int {
	hash := make(map[int]struct{})
	for _, v := range nums {
		hash[v] = struct{}{}
	}
	for i := 0; i <= len(nums); i++ {
		if _, ok := hash[i]; !ok {
			return i
		}
	}
	panic("")
}

func missingNumberSums(nums []int) int {
	numSum := 0
	for _, v := range nums {
		numSum += v
	}
	sum := 0
	for i := 0; i <= len(nums); i++ {
		sum += i
	}
	return sum - numSum
}
