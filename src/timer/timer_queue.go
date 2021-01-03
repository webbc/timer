package timer

type TimerQueue []*Timer

func (pq TimerQueue) Len() int { return len(pq) }

func (pq TimerQueue) Less(i, j int) bool {
	return pq[i].endTime < pq[j].endTime
}

func (pq TimerQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *TimerQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Timer)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *TimerQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0: n-1]
	return item
}
