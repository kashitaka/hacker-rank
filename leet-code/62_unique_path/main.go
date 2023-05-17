package _2_unique_path

func uniquePaths(m int, n int) int {
	routeCache := make([][]int, m)
	for i:=0; i<m; i++ {
		routeCache[i] = make([]int, n)
	}
	routeCache[0][0] = 1
	for i:=0; i<m;i++ {
		for j:=0; j<n; j++ {
			if i == 0 && j ==0 {
				continue
			}
			left, top := 0,0
			if j!=0 {
				left = routeCache[i][j-1]
			}
			if i!=0 {
				top = routeCache[i-1][j]
			}
			routeCache[i][j] = left+top
		}
	}
	return routeCache[m-1][n-1]
}