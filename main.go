package main

import (
	"fmt"
)
var cols[13] int
func dfs (n *int, row int ,ans *int) {
	if row == *n { //origin n = 12
		*ans += 1
	} else {
		for i := 1 ; i <= *n ; i++ {
			cols[ row + 1] = i
			if isPossible(row + 1) {
				dfs( n , row + 1, ans)
			} else {
				cols[row+1] = 0
			}
		}
	}
	cols[row] = 0
}
func isPossible (c int) bool {
	for i := 1 ; i < c ; i++ {
		if cols[i] == cols[c] {
			return false
		}
		if abs(cols[i] - cols[c]) == abs(i-c) {
			return false
		}
	}
	return true
}
func abs(num int) int {
	if num < 0 {
		return -1 * num
	} else {
		return num
	}
}
func solution(n int) int {
	ans := 0
	for i := 1 ; i <= n ; i++ {
		for j := 0 ; j <=12 ; j++ {
			cols[j] = 0
		}
		cols[1] = i
		dfs(&n,1,&ans)
	}
	return ans
}
func main() {
	fmt.Println(solution(7))
}