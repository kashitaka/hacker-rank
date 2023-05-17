package _09_fibonacci_number

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	fibRes := make(map[int]int)
	fibRes[0] = 0
	fibRes[1] = 1
	for i := 2; i <= n; i++ {
		fibRes[i] = fibRes[i-1] + fibRes[i-2]
	}
	return fibRes[n]
}