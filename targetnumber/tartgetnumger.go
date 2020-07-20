package targetnumber
var ans int = 0
func Solution(numbers []int, target int) int {
	dfs(0,0, numbers,target)
	return ans
}
func dfs (idx int, sum int,numbers[]int,target int){
	if idx == len(numbers) {
		if sum == target {
			ans += 1
		}
		return

	}
	dfs(idx+1, sum + numbers[idx], numbers, target)
	dfs(idx+1, sum - numbers[idx], numbers, target)
}