package _02_happy_number

func isHappy(n int) bool {
	hash := make(map[int]bool)

	for !hash[n] {
		digits := []int{}
		hash[n] = true
		for n != 0 {
			digits = append(digits, n%10)
			n = n / 10
		}
		newN := 0
		for _, v := range digits {
			newN += v * v
		}
		if newN == 1 {
			return true
		}
		n = newN
	}
	return false
}
