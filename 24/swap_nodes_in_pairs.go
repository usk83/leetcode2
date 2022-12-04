// 24. Swap Nodes in Pairs
// https://leetcode.com/problems/swap-nodes-in-pairs/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println("No local tests available.")
	pp.Println("=========================================")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * v2. Reconnect three nodes at once
 * - Time Complexity: O(N)
 * - Space Complexity: O(1)
 * > Runtime: 0 ms, faster than 100.00%
 * > Memory Usage: 2.1 MB, less than 15.13% (Probably the best mark)
 */
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	for p := dummy; p.Next != nil && p.Next.Next != nil; {
		p, p.Next, p.Next.Next, p.Next.Next.Next = p.Next, p.Next.Next, p.Next.Next.Next, p.Next
	}
	return dummy.Next
}

/**
 * v1. Prepare three pointers and re-connect nodes step by step
 * - Time Complexity: O(N)
 * - Space Complexity: O(1)
 * > Runtime: 0 ms, faster than 100.00%
 * > Memory Usage: 2.1 MB, less than 15.13%
 */
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	pointer := dummy
	for pointer.Next != nil && pointer.Next.Next != nil {
		first, second, third := pointer.Next, pointer.Next.Next, pointer.Next.Next.Next
		pointer.Next = second
		pointer.Next.Next = first
		pointer.Next.Next.Next = third
		pointer = pointer.Next.Next
	}
	return dummy.Next
}

// func swapPairs(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
//
// 	first, second, third := head, head.Next, head.Next.Next
// 	head = second
// 	head.Next = first
// 	head.Next.Next = third
//
// 	p := head.Next
// 	for p.Next != nil && p.Next.Next != nil {
// 		first, second, third := p.Next, p.Next.Next, p.Next.Next.Next
// 		p.Next = second
// 		p.Next.Next = first
// 		p.Next.Next.Next = third
// 		p = p.Next.Next
// 	}
//
// 	return head
// }
