package _071_greatest_common_divisor_of_string

func gcdOfStrings(str1 string, str2 string) string {
	if str1 + str2 != str2 + str1 {
		return ""
	}
	return str1[0:gcd(len(str1), len(str2))]
}

func gcd(m, n int) int {
	var min int
	if m > n {
		min = n
	} else {
		min = m
	}
	for min > 0 {
		if m % min ==0 && n % min ==0 {
			return min
		}
		min--
	}
	return min
}
