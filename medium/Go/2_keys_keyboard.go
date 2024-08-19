// link to the problem: https://leetcode.com/problems/2-keys-keyboard/

package main

import "fmt"

func minSteps(n int) int {
	if n == 1 {
			return 0
	}

	dp := make([]int, n+1)

	for i := 2; i <= n; i++ {
			dp[i] = i
			for j := i/2; j > 0; j--{
					if i%j == 0{
							dp[i] = dp[j] + (i/j)
							break
					}
			}
	} 
	return dp[n]
}

func main() {
	n := 3
	fmt.Println(minSteps(n)) // 3
}
