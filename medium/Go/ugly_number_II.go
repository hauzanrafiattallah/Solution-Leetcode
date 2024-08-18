// link to the problem: https://leetcode.com/problems/ugly-number-ii/

package main

import "math"

func nthUglyNumber(n int) int {
	t := make([]int, n+1)

	t[1] = 1
	t2 := 1
	t3 := 1
	t5 := 1

	for i := 2; i <= n; i++ {
		second := t[t2] * 2
		third := t[t3] * 3
		fifth := t[t5] * 5
		t[i] = int(math.Min(float64(second), math.Min(float64(third), float64(fifth))))

		if t[i] == second {
			t2++
		}
		if t[i] == third {
			t3++
		}
		if t[i] == fifth {
			t5++
		}
	}
	return t[n]
}

func main() {
	n := 10
	println(nthUglyNumber(n)) // 12
}
