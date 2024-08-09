// link to the problem: https://leetcode.com/problems/magic-squares-in-grid/

package main

import "fmt"


func isMagicSquare(grid [][]int, row int, col int) bool {
	// Magic square harus mengandung angka dari 1 hingga 9
	
	nums := make([]bool, 10)

	// Cek apakah semua angka dari 1 hingga 9 ada dalam subgrid
	for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
					num := grid[row+i][col+j]
					if num < 1 || num > 9 || nums[num] {
							return false
					}
					nums[num] = true
			}
	}

	// Hitung sum dari setiap baris, kolom, dan diagonal
	sum1 := grid[row][col] + grid[row][col+1] + grid[row][col+2]
	sum2 := grid[row+1][col] + grid[row+1][col+1] + grid[row+1][col+2]
	sum3 := grid[row+2][col] + grid[row+2][col+1] + grid[row+2][col+2]

	if sum1 != sum2 || sum2 != sum3 {
			return false
	}

	colSum1 := grid[row][col] + grid[row+1][col] + grid[row+2][col]
	colSum2 := grid[row][col+1] + grid[row+1][col+1] + grid[row+2][col+1]
	colSum3 := grid[row][col+2] + grid[row+1][col+2] + grid[row+2][col+2]

	if colSum1 != colSum2 || colSum2 != colSum3 {
			return false
	}

	diagSum1 := grid[row][col] + grid[row+1][col+1] + grid[row+2][col+2]
	diagSum2 := grid[row][col+2] + grid[row+1][col+1] + grid[row+2][col]

	if diagSum1 != diagSum2 || diagSum1 != sum1 {
			return false
	}

	return true
}

func numMagicSquaresInside(grid [][]int) int {
	count := 0

	// Iterasi melalui semua kemungkinan subgrid 3x3
	for i := 0; i < len(grid)-2; i++ {
			for j := 0; j < len(grid[0])-2; j++ {
					if isMagicSquare(grid, i, j) {
							count++
					}
			}
	}

	return count
}


func main() {
	grid := [][]int{{4,3,8,4},{9,5,1,9},{2,7,6,2}}
	fmt.Println(numMagicSquaresInside(grid)) // 1
}
