package main

import (
	"container/heap"
	"fmt"
	"math"
)

type City struct {
	x, y int
}

type Edge struct {
	cost float64
	to   int
}

type MinHeap []Edge

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Edge))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func euclideanDistance(c1, c2 City) float64 {
	return math.Sqrt(float64((c1.x-c2.x)*(c1.x-c2.x) + (c1.y-c2.y)*(c1.y-c2.y)))
}

func prim(n int, cities []City) float64 {
	visited := make([]bool, n)
	minHeap := &MinHeap{}
	heap.Push(minHeap, Edge{cost: 0, to: 0})
	maxEdge := 0.0

	for minHeap.Len() > 0 {
		edge := heap.Pop(minHeap).(Edge)
		cost, u := edge.cost, edge.to

		if visited[u] {
			continue
		}

		visited[u] = true
		if cost > maxEdge {
			maxEdge = cost
		}

		for v := 0; v < n; v++ {
			if !visited[v] {
				distance := euclideanDistance(cities[u], cities[v])
				heap.Push(minHeap, Edge{cost: distance, to: v})
			}
		}
	}

	return maxEdge / 2.0 // 同时施工，时间减半
}

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	cities := make([]City, n)

	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&cities[i].x, &cities[i].y)
	}

	result := prim(n, cities)
	fmt.Printf("%d\n", int(math.Ceil(result)))
}
