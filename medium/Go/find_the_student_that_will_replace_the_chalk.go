// link to problem: https://leetcode.com/problems/find-the-student-that-will-replace-the-chalk/

package main

import "fmt"

func chalkReplacer(chalk []int, k int) int {
	totalChalkNeeded := 0

	for _, studentChalkUse := range chalk {
		totalChalkNeeded += studentChalkUse
	}

	remainingChalk := k % totalChalkNeeded

	for studentIndex, studentChalkUse := range chalk {
		if remainingChalk < studentChalkUse {
			return studentIndex
		}
		remainingChalk -= studentChalkUse
	}

	return 0
}

func main() {
	// test case 1
	chalk := []int{5, 1, 5}
	k := 22
	fmt.Println(chalkReplacer(chalk, k)) // expected: 0
	// test case 2
	chalk = []int{3, 4, 1, 2}
	k = 25
	fmt.Println(chalkReplacer(chalk, k)) // expected: 1
}
