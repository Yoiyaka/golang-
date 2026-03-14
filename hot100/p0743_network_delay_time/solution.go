package p0743_network_delay_time

import (
	"container/heap"
	"math"
)

type Edge struct {
	to, weight int
}

type Item struct {
	node, dist int
}

type PQ []Item

func (pq PQ) Len() int            { return len(pq) }
func (pq PQ) Less(i, j int) bool  { return pq[i].dist < pq[j].dist }
func (pq PQ) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PQ) Push(x interface{}) { *pq = append(*pq, x.(Item)) }
func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

func networkDelayTime(times [][]int, n int, k int) int {
	adj := make([][]Edge, n+1)
	for _, t := range times {
		adj[t[0]] = append(adj[t[0]], Edge{t[1], t[2]})
	}
	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[k] = 0
	pq := &PQ{{k, 0}}
	heap.Init(pq)
	for pq.Len() > 0 {
		curr := heap.Pop(pq).(Item)
		if curr.dist > dist[curr.node] {
			continue
		}
		for _, e := range adj[curr.node] {
			newDist := dist[curr.node] + e.weight
			if newDist < dist[e.to] {
				dist[e.to] = newDist
				heap.Push(pq, Item{e.to, newDist})
			}
		}
	}
	maxDist := 0
	for i := 1; i <= n; i++ {
		if dist[i] == math.MaxInt32 {
			return -1
		}
		if dist[i] > maxDist {
			maxDist = dist[i]
		}
	}
	return maxDist
}
