package jump

func Solution(n int) int64 {
	memo := make([]int, n+2)
	memo[1] = 1
	memo[2] = 2
	for i := 3 ; i <= n ; i++ {
		memo[i] = ( memo[i-1] + memo[i-2] ) % 1234567
		// memo[n-1]칸에서 1칸 뛰어서 오는 방법 + memo[n-2]칸에서 2칸 뛰어서 오는 방법
	}
	/* 온라인 블로그들을 살펴보면 그냥 수열 n에 따른 수열 자체가 피보나치 수열을 따라서
	피보나치 수열 점화식으로 풀었다하는 글들이 많은데,
	위와 같이 DP 정의를 내려두는 것이 다른 문제풀 때도 도움이 될 것.
	 */
	return int64(memo[n])
}