// link to the problem: https://leetcode.com/problems/insert-greatest-common-divisors-in-linked-list/?envType=daily-question&envId=2024-09-10
package main

import "fmt"

// Definisikan struktur untuk node pada Linked List
type ListNode struct {
	Val  int
	Next *ListNode
}

// Fungsi untuk menghitung GCD (Greatest Common Divisor)
func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

// Fungsi untuk menambahkan GCD antara node yang berdekatan
func insertGCDs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head // Tidak ada node yang berdekatan
	}

	curr := head

	for curr != nil && curr.Next != nil {
		gcdVal := GCD(curr.Val, curr.Next.Val) // Hitung GCD antara dua node
		newNode := &ListNode{Val: gcdVal}      // Buat node baru dengan nilai GCD
		newNode.Next = curr.Next               // Hubungkan node baru ke node berikutnya
		curr.Next = newNode                    // Hubungkan node saat ini ke node baru
		curr = newNode.Next                    // Lanjutkan ke node setelah node yang baru
	}

	return head
}

// Fungsi untuk menampilkan linked list (untuk debugging)
func printList(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Printf("%d ", curr.Val)
		curr = curr.Next
	}
	fmt.Println()
}

// Main function untuk menguji kode
func main() {
	// Contoh linked list: [18, 6, 10, 3]
	head := &ListNode{Val: 18}
	head.Next = &ListNode{Val: 6}
	head.Next.Next = &ListNode{Val: 10}
	head.Next.Next.Next = &ListNode{Val: 3}

	fmt.Println("Original List:")
	printList(head) // Tampilkan linked list sebelum modifikasi

	// Modifikasi linked list dengan menyisipkan node GCD
	head = insertGCDs(head)

	fmt.Println("Modified List:")
	printList(head) // Tampilkan linked list setelah modifikasi
}
