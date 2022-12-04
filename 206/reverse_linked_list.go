// 206. Reverse Linked List
// https://leetcode.com/problems/reverse-linked-list/
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	printListNodes(head)
	head = reverseList(head)
	printListNodes(head)
	pp.Println("=========================================")
}

func printListNodes(cur *ListNode) {
	vals := []int{}
	for cur != nil {
		vals = append(vals, cur.Val)
		cur = cur.Next
	}
	fmt.Println(vals)
}

/**
 * Constraints:
 * - The number of nodes in the list is the range [0, 5000].
 * - -5000 <= Node.val <= 5000
 *
 * N: Number of the given list nodes
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 2. recursive solution
 * - Time Complexity: O(N)
 * - Space Complexity: O(N)
 * > Runtime: 0 ms (faster than 100.00%)
 * > Memory Usage: 3.1 MB (less than 8.35%)
 */
// func reverseList(head *ListNode) *ListNode {
// 	var reverse func(*ListNode, *ListNode) *ListNode
// 	reverse = func(prev, head *ListNode) *ListNode {
// 		if head == nil {
// 			return prev
// 		}
// 		next := head.Next
// 		head.Next = prev
// 		return reverse(head, next)
// 	}
// 	return reverse(nil, head)
// }

/**
 * 1-1. iterative solution oprimized
 * - Time Complexity: O(N)
 * - Space Complexity: O(1)
 * > Runtime: 0 ms (faster than 100.00%)
 * > Memory Usage: 2.6 MB (less than 100.00%)
 */
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next, prev = prev, head
		head = next
	}
	return prev
}

/**
 * 1. iterative solution
 * - Time Complexity: O(N)
 * - Space Complexity: O(1)
 * > Runtime: 0 ms (faster than 100.00%)
 * > Memory Usage: 2.6 MB (less than 100.00%)
 */
// func reverseList(head *ListNode) *ListNode {
// 	if head == nil {
// 		return nil
// 	}
// 	next := head.Next
// 	head.Next = nil
// 	for next != nil {
// 		prev := head
// 		head, next = next, next.Next
// 		head.Next = prev
// 	}
// 	return head
// }
