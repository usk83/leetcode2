// 21. Merge Two Sorted Lists
// https://leetcode.com/problems/merge-two-sorted-lists/
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
 * - The number of nodes in both lists is in the range [0, 50].
 * - -100 <= Node.val <= 100
 * - Both list1 and list2 are sorted in non-decreasing order.
 *
 * N: Number of nodes in the list1
 * M: Number of nodes in the list2
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 3. Call itself Recursively (better)
 * - Time Complexity: O(N+M)
 * - Space Complexity: O(N+M)
 * > Runtime: 0 ms (faster than 100.00%)
 * > Memory Usage: 2.6 MB (less than 60.48%)
 *
 * Inspired by other's solution
 */
// func mergeTwoLists(l1, l2 *ListNode) *ListNode {
// 	if l1 == nil {
// 		return l2
// 	}
// 	if l2 == nil {
// 		return l1
// 	}
// 	if l1.Val < l2.Val {
// 		l1.Next = mergeTwoLists(l1.Next, l2)
// 		return l1
// 	} else { // l1.Val >= l2.Val
// 		l2.Next = mergeTwoLists(l1, l2.Next)
// 		return l2
// 	}
// }

/**
 * 2. Recursive way (poor)
 * - Time Complexity: O(N+M)
 * - Space Complexity: O(N+M)
 * > Runtime: 0 ms (faster than 100.00%)
 * > Memory Usage: 2.6 MB (less than 60.48%)
 */
// func mergeTwoLists(l1, l2 *ListNode) *ListNode {
// 	var recur func(head, l1, l2 *ListNode)
// 	recur = func(head, l1, l2 *ListNode) {
// 		if l1 == nil {
// 			head.Next = l2
// 			return
// 		}
// 		if l2 == nil {
// 			head.Next = l1
// 			return
// 		}
// 		if l1.Val < l2.Val {
// 			head.Next = l1
// 			recur(head.Next, l1.Next, l2)
// 		} else { // l1.Val >= l2.Val
// 			head.Next = l2
// 			recur(head.Next, l1, l2.Next)
// 		}
// 	}
// 	dummy := &ListNode{Val: -1}
// 	recur(dummy, l1, l2)
// 	return dummy.Next
// }

/**
 * 1. Iterative way
 * - Time Complexity: O(N+M)
 * - Space Complexity: O(N+M)
 * > Runtime: 0 ms (faster than 1000.00%)
 * > Memory Usage: 2.6 MB (less than 1000.00%)
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	current := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else { // l1.Val >= l2.Val
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}
	if l1 != nil {
		current.Next = l1
	} else if l2 != nil {
		current.Next = l2
	}
	return dummy.Next
}
