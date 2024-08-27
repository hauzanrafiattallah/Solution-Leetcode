// link to the problem: https://leetcode.com/problems/path-with-maximum-probability/
package main

import (
	"container/heap"
	"fmt"
)

type Edge struct {
	node int
	prob float64
}

type MaxHeap []Edge

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i].prob > h[j].prob }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(Edge)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func maxProbability(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {
	graph := make([][]Edge, n)
	for i, edge := range edges {
		a, b := edge[0], edge[1]
		prob := succProb[i]
		graph[a] = append(graph[a], Edge{b, prob})
		graph[b] = append(graph[b], Edge{a, prob})
	}

	probs := make([]float64, n)
	for i := range probs {
		probs[i] = 0.0
	}
	probs[start_node] = 1.0

	h := &MaxHeap{}
	heap.Push(h, Edge{start_node, 1.0})

	for h.Len() > 0 {
		cur := heap.Pop(h).(Edge)
		node, prob := cur.node, cur.prob

		if node == end_node {
			return prob
		}

		for _, neighbor := range graph[node] {
			newProb := prob * neighbor.prob
			if newProb > probs[neighbor.node] {
				probs[neighbor.node] = newProb
				heap.Push(h, Edge{neighbor.node, newProb})
			}
		}
	}

	return 0.0
}

func main() {
	n := 3
	edges := [][]int{{0, 1}, {1, 2}, {0, 2}}
	succProb := []float64{0.5, 0.5, 0.2}
	start := 0
	end := 2
	result := maxProbability(n, edges, succProb, start, end)
	fmt.Println(result) // Output: 0.25
}
