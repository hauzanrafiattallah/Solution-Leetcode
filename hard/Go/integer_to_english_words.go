// Description: https://leetcode.com/problems/integer-to-english-words/

package main

import (
	"fmt"
	"strings"
)

// Convert a non-negative integer num to its English words representation.
var (
	lessThan20 = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	tens       = []string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
	thousands  = []string{"", "Thousand", "Million", "Billion"}
)

// numberToWords function
func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	result := ""
	for i := 0; num > 0; i++ {
		if num%1000 != 0 {
			result = helper(num%1000) + thousands[i] + " " + result
		}
		num /= 1000
	}

	return strings.TrimSpace(result)
}

// helper function
func helper(num int) string {
	if num == 0 {
		return ""
	} else if num < 20 {
		return lessThan20[num] + " "
	} else if num < 100 {
		return tens[num/10] + " " + helper(num%10)
	} else {
		return lessThan20[num/100] + " Hundred " + helper(num%100)
	}
}

// func main
func main() {
	// test cases
	fmt.Println(numberToWords(123))        // One Hundred Twenty Three
	fmt.Println(numberToWords(12345))      // Twelve Thousand Three Hundred Forty Five
	fmt.Println(numberToWords(1234567))    // One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven
	fmt.Println(numberToWords(1234567891)) // One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven Thousand Eight Hundred Ninety One
}
