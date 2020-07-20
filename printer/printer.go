package printer

import (
	"container/heap"
	"container/list"
	"fmt"
)
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
	return (*ih)[i].priority > (*ih)[j].priority
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
	//lookup	map[interface{}]*item

}
//Construct
func New() PriorityQueue{
	return PriorityQueue{
		itemHeap: &itemHeap{},
		//lookup: make(map[interface{}]*item),
	}
}

func (p *PriorityQueue) Len() int {
	return p.itemHeap.Len()
}

func (p *PriorityQueue) Insert(v interface{}, priority int){
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

func (p *PriorityQueue) Pop() (interface{}){
	if len(*p.itemHeap) == 0 {
		return nil
	}
	item := heap.Pop(p.itemHeap).(*item)
	//delete(p.lookup, item.value)
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

type Doc struct {
	index		int
	priority	int
}
func Solution(priorities []int, location int) int {
	answer := 0
	maxheap := New()
	queue := list.New()

	for i := 0 ; i < len(priorities) ; i++ {
		maxheap.Insert(priorities[i], priorities[i])
		a := new(Doc)
		a.index , a.priority = i, priorities[i]
		queue.PushBack(a)
		fmt.Println("in : ", a)
	}
	fmt.Println("queue len : ", queue.Len())


	for queue.Len() > 0 {
		e := queue.Front()
		prior := e.Value.(*Doc).priority
		fmt.Println(prior)
		if prior != maxheap.Front().(*item).priority {
			queue.PushBack(queue.Front())
			queue.Remove(queue.Front())
		}
		index := e.Value.(*Doc).index
		if index == location {
			answer += 1
			break
		} else {
			answer += 1
			maxheap.Pop()
			queue.Remove(e)
		}

	}
	return answer
}