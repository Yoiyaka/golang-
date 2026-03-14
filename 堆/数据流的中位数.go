package main

import "container/heap"

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
func (h MaxHeap) Top() int { return h[0] }

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
func (h MinHeap) Top() int { return h[0] }

type MedianFinder struct {
	low  MaxHeap
	high MinHeap
}

func Constructor() MedianFinder {
	mf := MedianFinder{
		low:  MaxHeap{},
		high: MinHeap{},
	}
	heap.Init(&mf.low)
	heap.Init(&mf.high)
	return mf
}

func (this *MedianFinder) AddNum(num int) {
	if this.low.Len() == 0 || num <= this.low.Top() {
		heap.Push(&this.low, num)
	} else {
		heap.Push(&this.high, num)
	}

	if this.low.Len() < this.high.Len() {
		heap.Push(&this.low, heap.Pop(&this.high))
	} else if this.low.Len() > this.high.Len()+1 {
		heap.Push(&this.high, heap.Pop(&this.low))
	}

	if this.low.Len() > 0 && this.high.Len() > 0 && this.low.Top() > this.high.Top() {
		a := heap.Pop(&this.low).(int)
		b := heap.Pop(&this.high).(int)
		heap.Push(&this.low, b)
		heap.Push(&this.high, a)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.low.Len() > this.high.Len() {
		return float64(this.low.Top())
	}

	return (float64(this.low.Top()) + float64(this.high.Top())) / 2.0
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
