package delivery

const INF int =   20000001
func solution(N int, road [][]int, k int) int {
	answer := 0
	var src, dst , c int
	cost := make([][]int, N+1)
	for i := 1 ; i <= N ; i++ {
		cost[i] = make([]int, N+1)
	}

	for i := 1 ; i <= N ; i++ {
		for j := 1 ; j <= N ; j++ {
			if i == j {
				cost[i][j] = 0
			} else {
				cost[i][j] = INF
			}
		}
	}

	for i := 0 ; i < len(road) ; i++ {
		src, dst, c = road[i][0], road[i][1], road[i][2]
		if c < cost[src][dst] || c < cost[dst][src] {
			cost[src][dst] = c
			cost[dst][src] = c
		}
	}

	for k := 1 ; k <= N ; k++ {
		for i := 1 ; i <= N ; i++ {
			for j := 1 ; j <= N ; j++ {
				if cost[i][j] > cost[i][k] + cost[k][j] {
					cost[i][j] = cost[i][k] + cost[k][j]
				}
				if cost[j][i] > cost[j][k] + cost[k][i] {
					cost[j][i] = cost[j][k] + cost[k][i]
				}
			}
		}
	}

	for i := 1 ; i <= N ; i++ {
		if cost[1][i] <= k {
			answer += 1
		}
	}

	return answer
}