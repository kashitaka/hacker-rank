package main

import (
	"container/heap"
	"sort"
)

type Room struct {
	end int
}

type PriorityQueue []Room

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq PriorityQueue) Less(i, j int) bool {
	// endが早いものが優先
	return pq[i].end < pq[j].end
}
func (pq *PriorityQueue) Push(x interface{}) {
	new := x.(Room)
	*pq = append(*pq, new)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func minMeetingRooms(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var pq PriorityQueue
	heap.Init(&pq)
	for i, v := range intervals {
		if i == 0 {
			newRoom := Room{
				end: v[1],
			}
			heap.Push(&pq, newRoom)
			continue
		}
		minRoom := pq[0]
		if minRoom.end <= v[0] {
			r := heap.Pop(&pq)
			ro := r.(Room)
			ro.end = v[1]
			heap.Push(&pq, ro)
		} else {
			newRoom := Room{
				end: v[1],
			}
			heap.Push(&pq, newRoom)
		}
	}
	return len(pq)
}

//func minMeetingRooms(intervals [][]int) int {
//	sort.Slice(intervals, func(i, j int) bool {
//		return intervals[i][0] < intervals[j][0]
//	})
//	var rooms []int
//	for _, v := range intervals {
//		res := findAvailableRoom(rooms, v[0])
//		if res == -1 {
//			// make new room
//			rooms = append(rooms, v[1])
//		} else {
//			rooms[res] = v[1]
//		}
//	}
//	return len(rooms)
//}
//
//func findAvailableRoom(rooms []int, start int) int {
//	for i, v := range rooms {
//		if v <= start {
//			return i
//		}
//	}
//	// no available room
//	return -1
//}
