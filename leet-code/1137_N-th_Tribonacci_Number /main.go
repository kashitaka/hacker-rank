package _137_N_th_Tribonacci_Number_

func tribonacci(n int) int {
	resMap := make(map[int]int)
	resMap[0] = 0
	resMap[1] = 1
	resMap[2] = 1
	if n <= 2 {
		return resMap[n]
	}
	for i:=3; i <=n; i++ {
		resMap[i] = resMap[i-3] + resMap[i-2] + resMap[i-1]
	}
	return resMap[n]
}