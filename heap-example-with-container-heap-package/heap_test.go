package main

import (
	"container/heap"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	var minQueue PriorityMinQueue
	heap.Init(&minQueue)
	heap.Push(&minQueue, Element{value: 50})
	fmt.Println(minQueue)
	heap.Push(&minQueue, Element{value: 30})
	fmt.Println(minQueue)
	heap.Push(&minQueue, Element{value: 20})
	fmt.Println(minQueue)
	heap.Push(&minQueue, Element{value: 10})
	fmt.Println(minQueue)
	heap.Pop(&minQueue)
	fmt.Println(minQueue)
	heap.Pop(&minQueue)
	fmt.Println(minQueue)
	heap.Pop(&minQueue)
	fmt.Println(minQueue)
	heap.Pop(&minQueue)
	fmt.Println(minQueue)

	var maxQueue PriorityMaxQueue
	heap.Init(&maxQueue)
	heap.Push(&maxQueue, Element{value: 10})
	fmt.Println(maxQueue)
	heap.Push(&maxQueue, Element{value: 20})
	fmt.Println(maxQueue)
	heap.Push(&maxQueue, Element{value: 30})
	fmt.Println(maxQueue)
	heap.Push(&maxQueue, Element{value: 40})
	fmt.Println(maxQueue)
	heap.Pop(&maxQueue)
	fmt.Println(maxQueue)
	heap.Pop(&maxQueue)
	fmt.Println(maxQueue)
	heap.Pop(&maxQueue)
	fmt.Println(maxQueue)
	heap.Pop(&maxQueue)
	fmt.Println(maxQueue)
}
