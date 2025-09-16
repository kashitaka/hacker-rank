package _0_pow_x_n

// divide and conquer

// 2^10 = (2^5) * (2^5)
// 2^11 = (2^5) * (2^6)
//      = (2^5) * (2^5) * 2

// Time: O(log n)
// Space: O(log n) for calling recursive function
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	if n > 0 {
		return pow(x, n)
	}
	return pow(1/x, -n)
}

func pow(x float64, n int) float64 {
	if n == 1 {
		return x
	}
	p := pow(x, n/2)
	if n%2 == 1 {
		return p * p * x
	}
	return p * p
}
