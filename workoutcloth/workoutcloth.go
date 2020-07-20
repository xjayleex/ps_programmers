package workoutcloth

func solution(n int, lost []int, reserve []int) int {
	chk := make([]int, n+1)
	for i := 1 ; i <= n ; i++ {
		chk[i] = 1
	}
	for i := 0 ; i < len(reserve) ; i++ {
		chk[reserve[i]] = 2
	}
	for i := 0 ; i < len(lost) ; i++ {
		if chk[lost[i]] == 2 {
			chk[lost[i]] = 1
		} else {
			chk[lost[i]] = 0
		}
	}
	for i := 1 ; i <= n ; i++ {
		if chk[i] == 0 {
			if chk[i-1] == 2 {
				chk[i-1], chk[i] = 1, 1
			} else if i+1 <= n && chk[i+1] == 2 {
				chk[i], chk[i+1] = 1, 1
			}
		}
	}
	cnt := 0
	for i := 1 ; i <= n ; i++ {
		if chk[i] >= 1 {
			cnt += 1;
		}
	}
	return cnt
}
