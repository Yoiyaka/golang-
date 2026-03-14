package p0295_find_median_from_data_stream

import "container/heap"

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type MedianFinder struct {
	lo *MaxHeap // max heap for lower half
	hi *MinHeap // min heap for upper half
}

func Constructor() MedianFinder {
	lo := &MaxHeap{}
	hi := &MinHeap{}
	heap.Init(lo)
	heap.Init(hi)
	return MedianFinder{lo: lo, hi: hi}
}

func (mf *MedianFinder) AddNum(num int) {
	heap.Push(mf.lo, num)
	heap.Push(mf.hi, heap.Pop(mf.lo))
	if mf.hi.Len() > mf.lo.Len() {
		heap.Push(mf.lo, heap.Pop(mf.hi))
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	if mf.lo.Len() > mf.hi.Len() {
		return float64((*mf.lo)[0])
	}
	return float64((*mf.lo)[0]+(*mf.hi)[0]) / 2.0
}
