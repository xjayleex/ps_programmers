package prior

import (
	"container/heap"
	"sort"
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
/*
type Pair struct {
	first int
	second int
}

type PairSlice []*Pair

func (p PairSlice) Len() int {
	return len(p)
}

func (p PairSlice) Less(i,j int) bool {
	return (p)[i].first < (p)[j].first
}

func (p PairSlice) Swap(i,j int) {
	(p)[i], (p)[j] = (p)[j] ,(p)[i]
}

type JSlice [][]int

func (j JSlice) Len() int {
	return len(j)
}
func (j JSlice) Less(i,k int) bool {
	return j[i][0] < j[k][0]
}
func (j *JSlice) Swap(i,k int) {
	(*j)[i], (*j)[k] = (*j)[k] , (*j)[i]
}

type Job struct{
	inTime int
	endTime int
	exeTime int
}*/

func Solution2 (jobs [][]int) int {
	sort.Slice(jobs[:],func(i,j int) bool {
		return jobs[i][0] < jobs[j][0]
	})

	idx, time, answer:= 0,0,0

	jobLen := len(jobs)

	pq := New()
	for idx < jobLen || !pq.IsEmpty() {
		if jobLen > idx && time >= jobs[idx][0] {
			pq.Insert(jobs[idx],jobs[idx][1])
			idx++
			continue
		}
		if !pq.IsEmpty(){
			tmp := pq.Pop().([]int)
			time += tmp[1]
			answer += time - tmp[0]
		}else {
			time = jobs[idx][0]
		}
	}
	return answer / jobLen
}

/*
func Solution(jobs [][]int) int {

	var inPairSlice PairSlice
	var sum int = 0

	for i := 0 ; i < len(jobs) ; i++{
		inPairSlice = append(inPairSlice,&Pair{jobs[i][0],jobs[i][1]})
	}

	sLen := inPairSlice.Len()
	sort.Sort(inPairSlice)

	pq := New()
	sIdx := 0
	f := inPairSlice[0].first
	for f == inPairSlice[sIdx].first {
		pq.Insert(Job{ inPairSlice[sIdx].first,
			inPairSlice[sIdx].first + inPairSlice[sIdx].second, inPairSlice[sIdx].second}, inPairSlice[sIdx].second)
		sIdx++
	}
	nowTime := 0
	for !((sIdx > sLen - 1 ) && pq.IsEmpty()) {
		if !pq.IsEmpty() {
			if pq.Front().(*item).value.(Job).endTime == nowTime {
				job := pq.Pop().(Job)
				sum = sum + job.endTime - job.inTime
				if !pq.IsEmpty() {
					tmp := pq.Pop().(Job)
					tmp.endTime = nowTime + tmp.exeTime
					pq.Insert(tmp,tmp.exeTime)
				}
			}
		}
		for sIdx <= sLen - 1 && inPairSlice[sIdx].first == nowTime {
			pq.Insert(Job{inPairSlice[sIdx].first, -1,inPairSlice[sIdx].second}, inPairSlice[sIdx].second)
			sIdx += 1
			if sIdx > sLen - 1 {
				break
			}
		}


		//finally
		nowTime += 1
	}
	return sum / sLen
}
*/