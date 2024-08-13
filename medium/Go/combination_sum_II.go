// link to the problem: https://leetcode.com/problems/combination-sum-ii/

package main

import "fmt"

func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	var comb []int

	quickSort(candidates, 0, len(candidates)-1) // sort the candidates with using quick sort algorithm
	findCombination(candidates, target, 0, comb, &res)

	return res
}

func findCombination(candidates []int, target int, start int, comb []int, res *[][]int) {
	if target == 0 {
		combination := make([]int, len(comb))
		copy(combination, comb)
		*res = append(*res, combination)
		return
	}

	for i := start; i < len(candidates); i++ {
		// Skip duplicates
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}

		// if the current candidate exceeds the target, break the loop
		if candidates[i] > target {
			break
		}

		// include the candidate and recursively find combinations with the remaining target
		comb = append(comb, candidates[i])
		findCombination(candidates, target-candidates[i], i+1, comb, res)
		// backtrack by removing the last element
		comb = comb[:len(comb)-1]
	}
}

// quicSort algorithm implementation
func quickSort(arr []int, low, high int) {
	if low < high {
		pivotIndex := partition(arr, low, high)
		quickSort(arr, low, pivotIndex-1)
		quickSort(arr, pivotIndex+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return i
}

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	result := combinationSum2(candidates, target) // [[1 1 6] [1 2 5] [1 7] [2 6]]
	fmt.Println(result)

	candidates = []int{2, 4, 6, 8, 10, 22, 22}
	target = 12
	result = combinationSum2(candidates, target) // [[2 4 6]] [[2, 10]] [[4 8]]
	fmt.Println(result)
}
