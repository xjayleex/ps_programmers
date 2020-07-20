package fibonaci

func Solution(n int) int {
	memo := make([]int,n+1)
	memo[0] = 0
	memo[1] = 1
	memo[2] = 1
	for i := 3 ; i < n+1 ; i++ {
		memo[i] = memo[i-1] + memo[i-2]  % 1234567
	}
	return memo[n]
}