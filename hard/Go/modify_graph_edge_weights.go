// link to problem: https://leetcode.com/problems/modify-graph-edge-weights/

package main

import (
	"container/heap"
	"math"
)

func modifiedGraphEdges(n int, edges [][]int, source int, destination int, target int) [][]int {
	adjacencyList := make([][]pair, n)
	for i := 0; i < len(edges); i++ {
		nodeA, nodeB := edges[i][0], edges[i][1]
		adjacencyList[nodeA] = append(adjacencyList[nodeA], pair{nodeB, i})
		adjacencyList[nodeB] = append(adjacencyList[nodeB], pair{nodeA, i})
	}

	distances := make([][2]int, n)
	for i := 0; i < n; i++ {
		if i != source {
			distances[i] = [2]int{math.MaxInt32, math.MaxInt32}
		}
	}

	runDijkstra(adjacencyList, edges, distances, source, 0, 0)
	difference := target - distances[destination][0]
	if difference < 0 {
		return [][]int{}
	}

	runDijkstra(adjacencyList, edges, distances, source, difference, 1)
	if distances[destination][1] < target {
		return [][]int{}
	}

	for i := range edges {
		if edges[i][2] == -1 {
			edges[i][2] = 1
		}
	}

	return edges
}

type pair struct {
	node, edgeIndex int
}

func runDijkstra(adjacencyList [][]pair, edges [][]int, distances [][2]int, source, difference, run int) {
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{value: source, priority: 0})
	distances[source][run] = 0

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Item)
		currentNode, currentDistance := current.value, current.priority

		if currentDistance > distances[currentNode][run] {
			continue
		}

		for _, neighbor := range adjacencyList[currentNode] {
			nextNode, edgeIndex := neighbor.node, neighbor.edgeIndex
			weight := edges[edgeIndex][2]
			if weight == -1 {
				weight = 1
			}

			if run == 1 && edges[edgeIndex][2] == -1 {
				newWeight := difference + distances[nextNode][0] - distances[currentNode][1]
				if newWeight > weight {
					edges[edgeIndex][2] = newWeight
					weight = newWeight
				}
			}

			if distances[nextNode][run] > distances[currentNode][run]+weight {
				distances[nextNode][run] = distances[currentNode][run] + weight
				heap.Push(pq, &Item{value: nextNode, priority: distances[nextNode][run]})
			}
		}
	}
}

type Item struct {
	value    int
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].priority < pq[j].priority }
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
