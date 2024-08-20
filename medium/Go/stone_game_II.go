// link to the problem: https://leetcode.com/problems/stone-game-ii/


func stoneGameII(piles []int) int {
	n := len(piles)
	// Sum[i] will store the sum of all stones from pile i to the end
	sum := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
			sum[i] = sum[i+1] + piles[i]
	}
	
	dp := make([][]int, n)
	for i := range dp {
			dp[i] = make([]int, n+1)
	}

	// Calculate dp[i][M] using the transition logic
	for i := n - 1; i >= 0; i-- {
			for M := 1; M <= n; M++ {
					for X := 1; X <= 2 * M && i + X <= n; X++ {
							if i + X < n {
									dp[i][M] = max(dp[i][M], sum[i] - dp[i+X][max(M, X)])
							} else {
									dp[i][M] = sum[i]  // If X exceeds the bounds, take all remaining piles
							}
					}
			}
	}
	
	// The answer is the maximum stones Alice can collect starting from the first pile with M=1
	return dp[0][1]
}

func max(a, b int) int {
	if a > b {
			return a
	}
	return b
}
