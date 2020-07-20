package carpet
func solution(brown int, yellow int) []int {
	brown = (brown - 4) / 2
	x := brown
	y := 0
	answer := make([]int,0)
	for  {
		y += 1
		x = brown - y
		if x * y == yellow {
			answer = append(answer, x+2)
			answer = append(answer, y+2)
			break
		}
	}
	return answer
}
