package validParenthesis

type Data struct {
	data [28]int
}
func (d *Data) setData (sl []int){
	for i := 0 ; i < len(sl) ; i++ {
		d.data[i] = sl[i]
	}
	for i := len(sl) ; i < 28 ; i++ {
		d.data[i] = 0
	}
}

func Solution(n int) int {
	inSlice := make([]int, 2 * n)
	mp := make(map[Data]bool)
	var data Data
	for i := 0 ; i < n ; i++ {
		inSlice[i] = 1
		inSlice[n+i] = -1
	}
	answer := 0
	for perm := range GeneratePermutations(inSlice){
		data.setData(perm)
		_, exists := mp[data]
		if exists {
			continue
		}
		mp[data] = true
		sum := 0
		answer += 1
		for i := 0 ; i < 2*n ; i++ {
			sum += perm[i]
			if sum < 0 {
				answer -= 1
				break
			}
		}
	}
	return answer
}

func GeneratePermutations(data []int) <-chan []int {
	c := make(chan []int)
	i := 0
	go func(c chan []int) {
		i++
		defer close(c)
		permutate(c, data)
	}(c)
	return c
}
func permutate(c chan []int, inputs []int){
	output := make([]int, len(inputs))
	copy(output, inputs)
	c <- output

	size := len(inputs)
	p := make([]int, size + 1)
	for i := 0; i < size + 1; i++ {
		p[i] = i;
	}
	for i := 1; i < size;{
		p[i]--
		j := 0
		if i % 2 == 1 {
			j = p[i]
		}
		inputs[i], inputs[j] = inputs[j], inputs[i]
		output := make([]int, len(inputs))
		copy(output, inputs)
		c <- output
		for i = 1; p[i] == 0; i++{
			p[i] = i
		}
	}
}