package _90_reverse_bits

func reverseBits(n int) int {
	res := 0
	for i := 0; i < 32; i++ {
		res = res << 1
		lowest := (n >> i) & 1
		res += lowest
	}
	return res
}
