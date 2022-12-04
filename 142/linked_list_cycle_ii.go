// 142. Linked List Cycle II
// https://leetcode.com/problems/linked-list-cycle-ii/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println("No local tests available.")
	pp.Println("=========================================")
}

/**
 * Constraints:
 * - The number of the nodes in the list is in the range [0, 10^4].
 * - -10^5 <= Node.val <= 10^5
 * - pos is -1 or a valid index in the linked-list.
 *
 * N: Number of ListNodes
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 1. Floyd's Cycle Detection Algorithm
 * - Time Complexity: O(N)
 * - Space Complexity: O(1)
 * > Runtime: 4 ms (faster than 94.64%)
 * > Memory Usage: 3.8 MB (less than 100.00%)
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head.Next, head.Next.Next
	for slow != fast && fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	if slow != fast {
		return nil
	}
	slow1, slow2 := head, slow
	for slow1 != slow2 {
		slow1, slow2 = slow1.Next, slow2.Next
	}
	return slow1
}
