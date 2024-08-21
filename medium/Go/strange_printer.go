// link to the problem: https://leetcode.com/problems/strange-printer/

package main

import "fmt"

func strangePrinter(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = n // Initialize with a large number
		}
	}

	for i := n - 1; i >= 0; i-- { // Loop backwards for proper subproblem solving
		dp[i][i] = 1 // Base case: single character
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i][j-1] // If the characters are the same
			} else {
				dp[i][j] = dp[i][j-1] + 1
				for k := i; k < j; k++ {
					dp[i][j] = min(dp[i][j], dp[i][k]+dp[k+1][j])
				}
			}
		}
	}

	return dp[0][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// test case
	s := "aaabbb"
	fmt.Println(strangePrinter(s)) // 2
	s = "aba"
	fmt.Println(strangePrinter(s)) // 2
	s = "abc"
	fmt.Println(strangePrinter(s)) // 3
}
