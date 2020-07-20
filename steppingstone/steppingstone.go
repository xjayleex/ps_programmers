package steppingstone

import (
	"sort"
)

func solution(distance int, rocks []int, n int) int {
	tmpSlice := []int {0,distance}
	rocks = append(rocks, tmpSlice ...)
	sort.Slice(rocks[:], func(i, j int) bool {
		return rocks[i] < rocks[j]
	})
	left, right := 1 , distance
	ans, mid := 0 , 0
	for left <= right {
		cnt, prev := 0, 0
		mid = (left + right) / 2
		for i := 1 ; i < len(rocks)  ; i++ {
			if rocks[i] - prev < mid {
				cnt += 1
			} else {
				prev = rocks[i]
			}
		}
		if cnt <= n {
			if mid > ans {
				ans = mid
			}
			left = mid + 1 //avail, make min value roughly
		} else {
			right = mid - 1 // unavailable, make min value tightly
		}
	}
	return ans
}