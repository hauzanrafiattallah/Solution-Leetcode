// link to the problem: https://leetcode.com/problems/count-sub-islands/

package main

import "fmt"

func f(i int, j int, pgrid1 *[][]int, pgrid2 *[][]int) bool {
	grid1 := *pgrid1
	grid2 := *pgrid2

	if i < 0 || i >= len(grid1) || j < 0 || j >= len(grid1[0]) || grid2[i][j] == 0 {
		return true
	}

	grid2[i][j] = 0

	ret := true
	ret = f(i+1, j, pgrid1, pgrid2) && ret
	ret = f(i-1, j, pgrid1, pgrid2) && ret
	ret = f(i, j+1, pgrid1, pgrid2) && ret
	ret = f(i, j-1, pgrid1, pgrid2) && ret

	if grid1[i][j] == 0 {
		return false
	}

	return ret
}

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	ans := 0

	for i := 0; i < len(grid1); i++ {
		for j := 0; j < len(grid1[i]); j++ {
			if grid1[i][j] == 0 || grid2[i][j] == 0 {
				continue
			}
			if f(i, j, &grid1, &grid2) {
				ans++
			}
		}
	}

	return ans
}

func main() {
	grid1 := [][]int{{1, 1, 1, 0, 0}, {0, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 1, 1}}
	grid2 := [][]int{{1, 1, 1, 0, 0}, {0, 0, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 1, 1}}
	fmt.Println(countSubIslands(grid1, grid2))
}
