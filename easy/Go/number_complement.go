// link to the problem: https://leetcode.com/problems/number-complement/

package main

import (
	"fmt"
	"strconv"
)

func findComplement(num int) int {
	binary := strconv.FormatInt(int64(num), 2)

	complement := ""
	for _, bit := range binary {
		if bit == '0' {
			complement += "1"
		} else if bit == '1' {
			complement += "0"
		}
	}

	result, _ := strconv.ParseInt(complement, 2, 64)

	return int(result)
}

func main() {
	// test case
	num := 5
	fmt.Println(findComplement(num)) // 2
	// more
	fmt.Println(findComplement(num)) // 0
}
