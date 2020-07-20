package immigration

import (
	"math"
	"sort"
)

func Solution(n int, times []int) int64 {
	answer := math.MaxInt64
	sort.Slice(times[:], func(i, j int) bool {
		return times[i] > times[j]
	})
	left := 0
	right := times[0] * n
	mid := 0
	for left <= right {
		available := 0
		mid = (left + right) / 2
		for i := 0 ; i < len(times) ; i++ {
			available +=  mid / times[i]
		}

		if n > available { // not available
			left = mid + 1
		} else { // available
			right = mid -1
			if mid < answer {
				answer = mid
			}
		}
	}
	return int64(answer)
}