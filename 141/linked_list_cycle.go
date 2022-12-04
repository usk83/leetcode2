// 141. Linked List Cycle
// https://leetcode.com/problems/linked-list-cycle/
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
 * Naive
 */
func hasCycle(head *ListNode) bool {
	visited := map[*ListNode]bool{} // set
	for head != nil {
		if visited[head] {
			return true
		}
		head = head.Next
		visited[head] = true
	}
	return false
}

/**
 * Naive 1
 */
var VISITED = 1000000007
func hasCycle(head *ListNode) bool {
	for head != nil {
		if head == VISITED {
			return true
		}
		head.Val = VISITED
		head = head.Next
	}
	return false
}

/**
 * v2
 */
func hasCycle(head *ListNode) bool {
	for  {


	if head == nil {
		return false
	}
	for slow, fast := head, head.Next; slow != fast; slow, fast = slow.Next, fast.Next.Next {
		if fast == nil || fast.Next == nil {
			return false
		}

	}
	return true
}

/**
 * v1
 */
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	for slow, fast := head, head.Next; slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow, fast = slow.Next, fast.Next.Next
	}
	return true
}
