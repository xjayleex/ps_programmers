package connectIsland

import "sort"

var parent[101] int
func Solution(n int, costs [][]int) int {
	ans := 0
	for i := 0 ; i < n ; i++ {
		parent[i] = i
	}
	sort.Slice(costs[:], func(i, j int) bool {
		return costs[i][2] < costs[j][2]
	})
	for i := 0 ; i < len(costs) ; i++ {
		if findRoot(costs[i][0]) == findRoot(costs[i][1]) {
			continue
		}
		union(costs[i][0], costs[i][1])
		ans += costs[i][2]
	}
	return ans
}

func findRoot(x int) int {
	if parent[x] == x {
		return x
	} else {
		parent[x] = findRoot(parent[x])
		return parent[x]
	}
}

func union(x int, y int) {
	x = findRoot(x)
	y = findRoot(y)
	parent[y] = x
}