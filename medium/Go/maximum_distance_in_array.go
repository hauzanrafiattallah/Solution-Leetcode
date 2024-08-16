// link to the problem: https://leetcode.com/problems/maximum-distance-in-arrays/

package main

func maxDistance(arrays [][]int) int {
	// Inisialisasi nilai maksimum dan minimum dari array pertama
	minVal := arrays[0][0]
	maxVal := arrays[0][len(arrays[0])-1]

	maxDistance := 0

	// iterasi di mulai dari array ke dua
	for i := 1; i < len(arrays); i++ {
		// Hitung jarak maksimum dengan menggunakan nilai minimum dan maksimum yang ditemukan sejauh ini
		maxDistance = max(maxDistance, abs(arrays[i][len(arrays[i])-1]-minVal))
		maxDistance = max(maxDistance, abs(maxVal-arrays[i][0]))

		// Update nilai minimum dan maksimum
		minVal = min(minVal, arrays[i][0])
		maxVal = max(maxVal, arrays[i][len(arrays[i])-1])
	}
	return maxDistance
}

// Fungsi untuk menghitung nilai mutlak
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// fungsi untuk menentukan minimum
func min(x, y int) int {
	if x > y {
		return y
	}

	return x
}

// fungsi untuk menentukan maksimum
func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func main() {
	// test case
	arrays := [][]int{{1, 2, 3}, {4, 5}, {1, 2, 3}}
	println(maxDistance(arrays)) // 4
	arrays = [][]int{{1}, {1}}
	println(maxDistance(arrays)) // 0
	arrays = [][]int{{1, 4}, {0, 5}}
	println(maxDistance(arrays)) // 4
}
