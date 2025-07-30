package _91_number_of_1_bits

func hammingWeight(n int) int {
	res := 0
	for n > 0 {
		res++
		n = n & (n - 1)
	}
	return res
}
