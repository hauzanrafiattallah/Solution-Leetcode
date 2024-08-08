// link to the problem: https://leetcode.com/problems/spiral-matrix-iii/
package main

import (
	"fmt"
)

func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	result := make([][]int, rows*cols)
	index := 0
	x, y := rStart, cStart

	result[index] = []int{x, y}
	index++

	step := 1
	for index < rows*cols {
		for i := 0; i < 4; i++ {
			dx, dy := directions[i][0], directions[i][1]
			for s := 0; s < step; s++ {
				x += dx
				y += dy
				if x >= 0 && x < rows && y >= 0 && y < cols {
					result[index] = []int{x, y}
					index++
				}
			}
			if i%2 == 1 {
				step++
			}
		}
	}

	return result
}

func main() {
	rows, cols := 1, 4
	rStart, cStart := 0, 0
	result := spiralMatrixIII(rows, cols, rStart, cStart)

	for _, pos := range result {
		fmt.Println(pos)
	}
}
