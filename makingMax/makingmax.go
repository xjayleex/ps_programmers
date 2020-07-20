package makingMax

import "strings"

func Solution(number string, k int) string {
	//retlen := len(number) - k
	var answer strings.Builder
	n := len(number)
	X := 0
	// n -  reftlen + 1 == n - n + k
	for pivot :=  k + 1 ; pivot <= n ; pivot++ {
		var max uint8 = 0
		for i := X ; i < pivot ; i++ {
			if number[i] > max {
				max = number[i]
				X = i + 1
			}
		}
		answer.WriteString(string(max))
	}
	return answer.String()
}