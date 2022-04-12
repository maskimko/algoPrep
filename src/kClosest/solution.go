package main

import (
	"container/heap"
	"math"
)

type point struct {
	x, y int
}

func (p point) distance() float32 {
	return float32(math.Sqrt(float64(p.x*p.x + p.y*p.y)))
}

type pq []point

func (p pq) Len() int {
	return len(p)
}

func (p pq) Less(i, j int) bool {
	return p[i].distance() < p[j].distance()
}

func (p pq) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *pq) Push(x interface{}) {
	*p = append(*p, x.(point))
}

func (p *pq) Pop() interface{} {
	c := (*p)[len(*p)-1]
	*p = (*p)[0 : len(*p)-1]
	return c
}

func kClosest(points []point, k int) []point {
	h := make(pq, 0)
	heap.Init(&h)
	for _, p := range points {
		heap.Push(&h, p)
	}
	var result []point
	for i := 0; i < k; i++ {
		result = append(result, heap.Pop(&h).(point))
	}
	return result
}
