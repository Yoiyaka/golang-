package main

// MaxHeap: 通过把 Less 写成 ">" 来实现大堆
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // 大堆关键
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push/Pop 用指针接收者，操作底层切片长度
func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

/*小堆
package main

import "container/heap"

// MaxHeap: 通过把 Less 写成 ">" 来实现大堆
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push/Pop 用指针接收者，操作底层切片长度
func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}





*/
