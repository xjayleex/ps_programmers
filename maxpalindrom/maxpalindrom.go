package maxpalindrom

import "sort"

func Solution(s string) int {
	slc1 := make([]int, len(s))
	slc2 := make([]int, len(s))
	strlen := len(s)
	for i := 0 ; i < strlen ; i++ {
		slc1[i] = 1
		for j := 1 ; true ; j++ {
			if i - j < 0 || i + j >= strlen {
				break
			}
			if s[i-j] == s[i+j] {
				slc1[i] += 2
			} else {
				break
			}
		}
	}

	for i := 0 ; i < strlen - 1 ; i++ {
		if s[i] != s[i+1] {
			continue
		}
		slc2[i] = 2
		for j := 1 ; true ; j++ {
			if i - j < 0 || i + j + 1 >= strlen {
				break
			}
			if s[i-j] == s[i+j+1] {
				slc2[i] += 2
			} else {
				break
			}
		}
	}

	sort.Slice(slc1[:], func(i, j int) bool {
		return slc1[i] > slc1[j]
	})
	sort.Slice(slc2[:], func(i, j int) bool {
		return slc2[i] > slc2[j]
	})
	ans := 0
	if slc1[0] > slc2[0] {
		ans = slc1[0]
	} else {
		ans = slc2[0]
	}
	return ans
}