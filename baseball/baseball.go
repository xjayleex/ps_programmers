package baseball

import (
	"strconv"
)

func Solution(baseball [][]int) int {
	answer := 0
	strikeCnt, ballCnt := 0, 0
	numStr :=""
	guess := ""
	for num := 100 ; num <= 999 ; num++ {
		numStr = strconv.Itoa(num)
		if !isValid(numStr) {
			continue
		}
		flag := true
		for i := 0 ; i < len(baseball) ; i++ {
			guess = strconv.Itoa(baseball[i][0])
			strikeCnt, ballCnt = check(numStr,guess)
			if strikeCnt == baseball[i][1] && ballCnt == baseball[i][2] {
				continue
			} else {
				flag = false
				break
			}
		}
		if flag == true {
			answer += 1
		}
	}
	return answer
}
func isValid(num string) bool {
	if num[0] == '0' || num[1] == '0' || num[2] == '0'{
		return false
	}
	if num[0] == num[1] {
		return false
	}
	if num[0] == num[2] {
		return false
	}
	if num[1] == num[2] {
		return false
	}
	return true
}
func check (origin string, guess string) (strikeCnt int, ballCnt int){
	// strike check
	for i := 0 ; i < 3 ; i++ {
		for j := 0 ; j < 3 ; j++ {
			if origin[j] == guess[i] {
				if i == j {
					strikeCnt += 1
				} else {
					ballCnt += 1
				}

			}
		}
	}
	return
}

