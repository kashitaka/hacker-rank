package _6_plus_one

func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		n := digits[i] + carry
		if n >= 10 {
			n = 0
			carry = 1
		} else {
			carry = 0
		}
		digits[i] = n
	}
	if carry == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
