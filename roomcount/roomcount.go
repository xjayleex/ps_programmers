package roomcount

const INF int = -987654321
type Tuple struct {
	first int
	second int
	third int
	forth int
}
func Solution(arrows []int) int {
	answer := 0
	dx := [8]int{0, 1, 1, 1, 0, -1, -1, -1}
	dy := [8]int{-1,-1,0, 1, 1,  1,  0, -1}
	vertexChk := make(map[Tuple]int)
	edgeChk := make(map[Tuple]int)
	vertexChk[Tuple{0,0, INF, INF}] = 1
	x, y := 0, 0
	for i := 0 ; i < len(arrows) ; i++ {
		for j := 0 ; j < 2 ; j++ {
			nx := x + dx[arrows[i]]
			ny := y + dy[arrows[i]]
			_, exist := vertexChk[Tuple{nx,ny,INF,INF}]
			if exist {
				if edgeChk[Tuple{nx,ny, x, y}] == 0 ||
					edgeChk[Tuple{x,y, nx, ny}] == 0 {
					answer += 1
				}
			}
			vertexChk[Tuple{nx,ny,INF,INF}] = 1
			edgeChk[Tuple{x,y,nx,ny}] = 1
			edgeChk[Tuple{nx,ny,x,y}] = 1

			x , y = nx , ny
		}
	}
	return answer
}