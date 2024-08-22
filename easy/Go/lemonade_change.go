// link to the problem: https://leetcode.com/problems/lemonade-change/

package main

import "fmt"


func lemonadeChange(bills []int) bool {
  five, ten := 0, 0
	for _, bill := range bills {
		switch bill {
		case 5:
			five++
		case 10:
			if five == 0{
				return false
			}
			five--
			ten++
		case 20:
			if ten > 0 && five > 0 {
				ten--
				five--
			}else if five >= 3 {
				five -= 3
			}else {
				return false
			}
		}
	}
	return true
}

func main()  {
	// test case
	fmt.Println(lemonadeChange([]int{5,5,5,10,20})) // true
	fmt.Println(lemonadeChange([]int{5,5,10})) // true
	fmt.Println(lemonadeChange([]int{10,10})) // false
}
