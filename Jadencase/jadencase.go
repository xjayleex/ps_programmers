package Jadencase

import "strings"

func Solution(s string) string {
	slower := strings.ToLower(s)
	var answer strings.Builder
	isPrevBlank := true
	for i := 0 ; i < len(slower) ;i++{
		if slower[i] == 32 {
			answer.WriteByte(32)
			isPrevBlank = true
		} else {
			if isPrevBlank {
				if slower[i] >= 97 && slower[i] <= 122 {
					answer.WriteByte(slower[i] - 32)
				} else {
					answer.WriteByte(slower[i])
				}
				isPrevBlank = false
			} else {
				answer.WriteByte(slower[i])
			}
		}
	}
	return answer.String()
}