package dobulepq

import (
	"container/heap"
	"strconv"
)

func Solution(operations []string) []int {
	answer := make([]int,0)
	var op uint8
	num := 0
	var chk [1000001]bool
	var lookup [1000001]int
	minHeap := New()
	maxHeap := New2()
	for i :=0 ; i < len(operations) ; i++ {
		op = operations[i][0] //73:I 68:D
		num, _= strconv.Atoi(string(operations[i][2:]))
		if op == 73 {
			minHeap.Insert(i,num) // value, priority
			maxHeap.Insert(i,num)
			lookup[i] = num
			chk[i] = true
		} else {
			if num == 1{
				for !maxHeap.IsEmpty2() && !chk[maxHeap.Front2().(*item).value.(int)] {
					maxHeap.Pop2()
				}
				if !maxHeap.IsEmpty2(){
					chk[maxHeap.Front2().(*item).value.(int)] = false
					maxHeap.Pop2()
				}
			} else if num == -1 {
				for !minHeap.IsEmpty() && !chk[minHeap.Front().(*item).value.(int)]{
					minHeap.Pop()
				}
				if !minHeap.IsEmpty(){
					chk[minHeap.Front().(*item).value.(int)] = false
					minHeap.Pop()
				}
			}
		}
	}
	for !maxHeap.IsEmpty2() && !chk[maxHeap.Front2().(*item).value.(int)] {
		maxHeap.Pop2()
	}
	for !minHeap.IsEmpty() && !chk[minHeap.Front().(*item).value.(int)]{
		minHeap.Pop()
	}

	if maxHeap.IsEmpty2() && minHeap.IsEmpty() {
		answer = append(answer, 0)
		answer = append(answer, 0)
	} else {
		answer = append(answer,lookup[maxHeap.Pop2().(int)])
		answer = append(answer,lookup[minHeap.Pop().(int)])
	}
	return answer
}

type item struct {
	value		interface{}
	priority	int
	index		int
}

type itemHeap []*item

func (ih *itemHeap) Len() int{
	return len(*ih)
}

func (ih *itemHeap) Less(i, j int) bool {
	return (*ih)[i].priority < (*ih)[j].priority
}

func (ih *itemHeap) Front() (interface{}){
	if len(*ih) == 0{
		return nil
	}
	return (*ih)[0]
}

func (ih *itemHeap) Swap(i, j int) {
	(*ih)[i], (*ih)[j] = (*ih)[j], (*ih)[i]
	(*ih)[i].index = i
	(*ih)[j].index = j
}

func (ih *itemHeap) Push(x interface{}){
	item := x.(*item)
	item.index = len(*ih)
	*ih = append(*ih,item)
}

func (ih *itemHeap) Pop()(x interface{}){
	old := *ih
	item := old[len(old)-1]
	*ih = old[0 : len(old)-1]
	return item
}

type PriorityQueue struct{
	itemHeap *itemHeap
	//lookup	map[interface{}]int

}
//Construct
func New() PriorityQueue{
	return PriorityQueue{
		itemHeap: &itemHeap{},
		//lookup: make(map[interface{}]int),
	}
}

func (p *PriorityQueue) Len() int {
	return p.itemHeap.Len()
}

func (p *PriorityQueue) Insert(v interface{}, priority int){
	/*cnt, exists := p.lookup[v]
	if exists {
		p.lookup[v] = cnt + 1
	} else {
		p.lookup[v] = 1
	}*/
	newItem := &item{
		value: v,
		priority: priority,
	}
	heap.Push(p.itemHeap, newItem)
}

func (p *PriorityQueue) Pop() (interface{}){
	if len(*p.itemHeap) == 0 {
		return nil
	}
	item := heap.Pop(p.itemHeap).(*item)

	/*cnt, _ := p.lookup[item.value]
	if cnt > 1 {
		p.lookup[item.value] = cnt - 1
	} else {
		delete(p.lookup, item.value)
	}*/

	return item.value
}


func (p *PriorityQueue) Front() (interface{}){
	if len(*p.itemHeap) == 0 {
		return nil
	}
	return p.itemHeap.Front().(*item)
}

func (p *PriorityQueue) IsEmpty() bool{
	if len(*p.itemHeap) == 0 {
		return true
	} else {
		return false
	}
}

/**************************/
type itemHeap2 []*item

func (ih *itemHeap2) Len() int{
	return len(*ih)
}

func (ih *itemHeap2) Less(i, j int) bool {
	return (*ih)[i].priority > (*ih)[j].priority
}

func (ih *itemHeap2) Front() (interface{}){
	if len(*ih) == 0{
		return nil
	}
	return (*ih)[0]
}

func (ih *itemHeap2) Swap(i, j int) {
	(*ih)[i], (*ih)[j] = (*ih)[j], (*ih)[i]
	(*ih)[i].index = i
	(*ih)[j].index = j
}

func (ih *itemHeap2) Push(x interface{}){
	item := x.(*item)
	item.index = len(*ih)
	*ih = append(*ih,item)
}

func (ih *itemHeap2) Pop()(x interface{}){
	old := *ih
	item := old[len(old)-1]
	*ih = old[0 : len(old)-1]
	return item
}

type PriorityQueue2 struct{
	itemHeap *itemHeap2
	//lookup	map[interface{}]*item

}
//Construct
func New2() PriorityQueue2{
	return PriorityQueue2{
		itemHeap: &itemHeap2{},
		//lookup: make(map[interface{}]*item),
	}
}

func (p *PriorityQueue2) Len() int {
	return p.itemHeap.Len()
}

func (p *PriorityQueue2) Insert(v interface{}, priority int){
	/*_, ok := p.lookup[v]
	if ok {
		return
	}*/
	newItem := &item{
		value: v,
		priority: priority,
	}
	heap.Push(p.itemHeap, newItem)
	//p.lookup[v] = newItem
}

func (p *PriorityQueue2) Pop2() (interface{}){
	if len(*p.itemHeap) == 0 {
		return nil
	}
	item := heap.Pop(p.itemHeap).(*item)
	//delete(p.lookup, item.value)
	return item.value
}


func (p *PriorityQueue2) Front2() (interface{}){
	if len(*p.itemHeap) == 0 {
		return nil
	}
	return p.itemHeap.Front().(*item)
}

func (p *PriorityQueue2) IsEmpty2() bool{
	if len(*p.itemHeap) == 0 {
		return true
	} else {
		return false
	}
}