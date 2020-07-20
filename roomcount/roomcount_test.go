package roomcount

import "fmt"

func main() {
	inSlice := [] int{
		6, 6, 6, 4, 4, 4, 2, 2, 2, 0, 0, 0, 1, 6, 5, 5, 3, 6, 0,
	}
	fmt.Println(Solution(inSlice))
}
