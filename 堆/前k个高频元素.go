package main

import "container/heap"

type Item struct {
	val  int
	freq int
}

// 大堆：freq 越大优先级越高
type MaxHeap []Item

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
	if h[i].freq != h[j].freq {
		return h[i].freq > h[j].freq
	}
	// 频率相同随便；加一个稳定的次序避免不确定性也可以
	return h[i].val > h[j].val
}
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) { *h = append(*h, x.(Item)) }
func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	cnt := make(map[int]int, len(nums))
	for _, x := range nums {
		cnt[x]++
	}

	h := make(MaxHeap, 0, len(cnt))
	for v, f := range cnt {
		h = append(h, Item{val: v, freq: f})
	}
	heap.Init(&h)

	res := make([]int, 0, k)
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(&h).(Item).val)
	}
	return res
}
