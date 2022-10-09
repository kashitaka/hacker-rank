package main

type Element struct {
	value int
}

type PriorityMinQueue []Element

func (pq PriorityMinQueue) Len() int {
	return len(pq)
}
func (pq PriorityMinQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq PriorityMinQueue) Less(i, j int) bool {
	// valが小さいものが優先
	return pq[i].value < pq[j].value
}
func (pq *PriorityMinQueue) Push(x interface{}) {
	new := x.(Element)
	*pq = append(*pq, new)
}
func (pq *PriorityMinQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

type PriorityMaxQueue []Element

func (pq PriorityMaxQueue) Len() int {
	return len(pq)
}
func (pq PriorityMaxQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq PriorityMaxQueue) Less(i, j int) bool {
	// valが大きいものが優先
	return pq[i].value > pq[j].value
}
func (pq *PriorityMaxQueue) Push(x interface{}) {
	new := x.(Element)
	*pq = append(*pq, new)
}
func (pq *PriorityMaxQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
