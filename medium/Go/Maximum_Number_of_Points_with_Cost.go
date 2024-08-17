// link to the problem: https://leetcode.com/problems/maximum-number-of-points-with-cost/

package main

import (
	"fmt"
	"math"
)

func maxPoints(points [][]int) int64 {
	m, n := len(points), len(points[0])

	dp := make([]int64, n)
	for j := 0; j < n; j++ {
		dp[j] = int64(points[0][j])
	}

	for i := 1; i < m; i++ {
		left := make([]int64, n)
		right := make([]int64, n)

		left[0] = dp[0]
		for j := 1; j < n; j++ {
			left[j] = max(left[j-1]-1, dp[j])
		}

		right[n-1] = dp[n-1]
		for j := n - 2; j >= 0; j-- {
			right[j] = max(right[j+1]-1, dp[j])
		}

		for j := 0; j < n; j++ {
			dp[j] = int64(points[i][j]) + max(left[j], right[j])
		}
	}

	maxPoints := int64(math.MinInt64)
	for _, val := range dp {
		maxPoints = max(maxPoints, val)
	}

	return maxPoints
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	points := [][]int{{1, 2, 3}, {1, 5, 1}, {3, 1, 1}}
	fmt.Println(maxPoints(points)) // 9
	points = [][]int{{1, 5}, {2, 3}, {4, 2}}
	fmt.Println(maxPoints(points)) // 11
}
