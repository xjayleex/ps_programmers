package hanoi

//move(from,to,temp)
//move(1,3,2)
var result [][] int
func move(n int, from int, to int, temp int){
	if n == 1 {
		t := make([]int,2)
		t[0], t[1] = from, to
		result = append(result,t)
	} else {
		move(n-1, from, temp, to)
		t := make([]int,2)
		t[0], t[1] = from, to
		result = append(result,t)
		move(n-1, temp, to, from)
	}
}
func Solution(n int) [][]int {
	move(n,1,3,2)
	return result
}