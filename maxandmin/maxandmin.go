package maxandmin

import (
	"strconv"
	"strings"
)

func Solution(s string) string {
	splited := strings.Split(s," ")
	min := 987654321
	max := -987654321
	for i := 0 ; i < len(splited) ; i++ {
		now, _ := strconv.Atoi(splited[i])
		if now > max {
			max = now
		}
		if now < min {
			min = now
		}
	}
	answer :=strconv.Itoa(min) + " "+ strconv.Itoa(max)
	return answer
}