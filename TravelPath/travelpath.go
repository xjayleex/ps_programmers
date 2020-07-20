package travelPath

import (
	"sort"
	"strings"
)

var chk [10001] bool
var answer [] string

func Solution_travelpath2(tickets [][]string) []string {
	ret := make([]string,0)
	answer = make([]string,0)
	ans := make([]string,0)
	sort.Slice(tickets[:], func(i,j int) bool{
		comp := strings.Compare(tickets[i][0],tickets[j][0])
		if comp == 0{
			return tickets[i][1] < tickets[j][1]
		} else if comp == -1{
			return true
		} else {
			return false
		}
	})
	dfs("ICN",0,len(tickets),tickets,ans)
	ret = append(ret, "ICN")
	for i := 0 ; i < len(answer) ; i++ {
		ret = append(ret, answer[i])

	}
	return ret
}

func dfs (start string, cnt int, threshold int,
	tickets [][] string , ans []string) bool{
	if cnt == threshold {

		for i := 0 ; i < len(ans) ; i++{
			answer = append(answer, ans[i])
		}
		return true
	}
	for i := 0 ; i < len(tickets) ; i++{
		if strings.Compare(start, tickets[i][0]) == 0 &&  !chk[i] {
			chk[i] = true
			ans = append(ans, tickets[i][1])
			result := dfs(tickets[i][1], cnt + 1, threshold, tickets, ans)
			if result{
				return true
			}
			chk[i] = false
			ans = ans[:len(ans)-1]
		}
	}
	return false
}

