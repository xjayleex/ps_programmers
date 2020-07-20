package ranking

func Solution(n int, results [][]int) int {
	answer := 0
	memo := make([][]bool, n+1 )
	for i:=0 ; i < n+1 ; i++ {
		memo[i] = make([]bool, n+1)
	}

	for i:=0 ; i < len(results) ;i++ {
		memo[results[i][0]][results[i][1]] = true
	}

	for k := 1 ; k <= n ; k++ {
		for i := 1 ; i <= n ; i++{
			for j := 1 ; j <= n ; j++ {
				if memo[i][k] && memo[k][j] {
					memo[i][j] = true
				}
			}
		}
	}

	for i := 1 ; i <= n ; i++ {
		cnt := 0
		for j := 1 ; j <= n ; j++ {
			if memo[i][j] || memo[j][i] {
				cnt += 1
			}
		}
		if cnt == n-1 {
			answer += 1
		}
	}
	return answer
}