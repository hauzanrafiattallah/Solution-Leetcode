// link to the problem: https://leetcode.com/problems/most-stones-removed-with-same-row-or-column/


package main

import "fmt"

type UnionFind struct {
	parent map[int]int
	rank   map[int]int
}

func NewUnionFind() *UnionFind {
	return &UnionFind{
			parent: make(map[int]int),
			rank:   make(map[int]int),
	}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
			uf.parent[x] = uf.Find(uf.parent[x]) // path compression
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX != rootY {
			if uf.rank[rootX] > uf.rank[rootY] {
					uf.parent[rootY] = rootX
			} else if uf.rank[rootX] < uf.rank[rootY] {
					uf.parent[rootX] = rootY
			} else {
					uf.parent[rootY] = rootX
					uf.rank[rootX]++
			}
	}
}

func (uf *UnionFind) Add(x int) {
	if _, exists := uf.parent[x]; !exists {
			uf.parent[x] = x
			uf.rank[x] = 0
	}
}

func removeStones(stones [][]int) int {
	uf := NewUnionFind()
	
	for _, stone := range stones {
			row, col := stone[0], stone[1] + 10001 // offset to differentiate rows and columns
			uf.Add(row)
			uf.Add(col)
			uf.Union(row, col)
	}
	
	// Count unique components
	componentSet := make(map[int]bool)
	for _, stone := range stones {
			row := stone[0]
			componentSet[uf.Find(row)] = true
	}
	
	return len(stones) - len(componentSet)
}

func main() {
	stones := [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}}
	fmt.Println(removeStones(stones))
}