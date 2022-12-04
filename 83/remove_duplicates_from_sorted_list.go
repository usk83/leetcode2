// 83. Remove Duplicates from Sorted List
// https://leetcode.com/problems/remove-duplicates-from-sorted-list/
package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	list1 := &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}}
	fmt.Println("Before list1:", ListNodesToString(list1))
	list2 := deleteDuplicates(list1)
	fmt.Println("After list1: ", ListNodesToString(list1))
	fmt.Println("After list2: ", ListNodesToString(list2))
	pp.Println("=========================================")
}

/**
 * Constraints:
 * - The number of nodes in the list is in the range [0, 300].
 * - -100 <= Node.val <= 100
 * - The list is guaranteed to be sorted in ascending order.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func ListNodesToString(node *ListNode) string {
	var nums []string
	for n := node; n != nil; n = n.Next {
		nums = append(nums, strconv.Itoa(n.Val))
	}
	return fmt.Sprintf("[%s]", strings.Join(nums, ", "))
}

/**
 * 1-1. Iterative, Modify the original list
 * - Time Complexity: O(N)
 * - Space Complexity: O(1)
 */

/**
 * 1-1-1. Shorter
 * > Runtime: 0 ms (faster than 100.00%)
 * > Memory Usage: 3 MB (less than 93.22%)
 */
// func deleteDuplicates(head *ListNode) *ListNode {
// 	node := head
// 	for node != nil && node.Next != nil {
// 		if node.Val == node.Next.Val {
// 			node.Next = node.Next.Next
// 		} else {
// 			node = node.Next
// 		}
// 	}
// 	return head
// }

/**
 * 1-1-2. Clearer
 * > Runtime: 4 ms (faster than 83.13%)
 * > Memory Usage: 3 MB (less than 93.22%)
 */
// func deleteDuplicates(head *ListNode) *ListNode {
// 	if head == nil {
// 		return nil
// 	}
// 	node := head
// 	for node.Next != nil {
// 		if node.Val == node.Next.Val {
// 			node.Next = node.Next.Next
// 		} else {
// 			node = node.Next
// 		}
// 	}
// 	return head
// }

/**
 * 1-2. Iterative, Create a new list
 * - Time Complexity: O(N)
 * - Space Complexity: O(N)
 */
// func deleteDuplicates(head *ListNode) *ListNode {
// 	if head == nil {
// 		return nil
// 	}
// 	newHead := &ListNode{head.Val, nil}
// 	newNode := newHead
// 	next := head.Next
// 	for next != nil {
// 		if head.Val == next.Val {
// 			next = next.Next
// 		} else {
// 			newNode.Next = &ListNode{next.Val, nil}
// 			newNode = newNode.Next
// 			head, next = next, next.Next
// 		}
// 	}
// 	return newHead
// }

/**
 * 2-1. recursive, Modify the original list
 * - Time Complexity: O(N)
 * - Space Complexity: O(N)
 */

/**
 * 2-1-1. Own solution
 */
// func deleteDuplicates(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	if head.Val == head.Next.Val {
// 		head.Next = head.Next.Next
// 		return deleteDuplicates(head)
// 	} else {
// 		head.Next = deleteDuplicates(head.Next)
// 		return head
// 	}
// }

/**
 * 2-1-2. Shorter and Clearer
 * ref: https://leetcode.com/problems/remove-duplicates-from-sorted-list/discuss/28625/3-Line-JAVA-recursive-solution
 */
// func deleteDuplicates(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	head.Next = deleteDuplicates(head.Next)
// 	if head.Val == head.Next.Val {
// 		return head.Next
// 	}
// 	return head
// }

/**
 * 2-2. recursive, Create a new list
 * - Time Complexity: O(N)
 * - Space Complexity: O(N)
 */
// func deleteDuplicates(head *ListNode) *ListNode {
// 	if head == nil {
// 		return nil
// 	}
// 	newNode := &ListNode{
// 		Val: head.Val,
// 	}
// 	if head.Next == nil {
// 		return newNode
// 	} else if head.Val != head.Next.Val {
// 		newNode.Next = deleteDuplicates(head.Next)
// 		return newNode
// 	} else {
// 		newNode.Next = head.Next.Next
// 		return deleteDuplicates(newNode)
// 	}
// }
