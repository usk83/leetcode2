// 111. Minimum Depth of Binary Tree
// https://leetcode.com/problems/minimum-depth-of-binary-tree/
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
 * - The number of nodes in the tree is in the range [0, 105].
 * - -1000 <= Node.val <= 1000
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type nodeItem struct {
	node *TreeNode
	rank int
}

/**
 * [description]
 * - Time Complexity: O(X)
 * - Space Complexity: O(X)
 * > Runtime: X ms (faster than 100.00%)
 * > Memory Usage: X MB (less than 100.00%)
 */

func minDepth(root *TreeNode) int {
	queue := []*TreeNode{&TreeNode{Left: root}}
	var depth int
LOOP:
	for len(queue) != 0 {
		nextQueue := []*TreeNode{}
		for _, node := range queue {
			if node == nil {
				continue
			}
			if node.Left == nil && node.Right == nil {
				break LOOP
			}
			if node.Left != nil {
				nextQueue = append(nextQueue, node.Left)
			}
			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right)
			}
		}
		depth++
		queue = nextQueue
	}
	return depth
}

// func minDepth(root *TreeNode) int {
// 	queue := []nodeItem{
// 		nodeItem{node: &TreeNode{Left: root}, rank: 0},
// 	}
// 	for len(queue) != 0 {
// 		item := queue[0]
// 		node := item.node
// 		if node == nil {
// 			queue = queue[1:]
// 			continue
// 		}
// 		if node.Left == nil && node.Right == nil {
// 			return item.rank
// 		}
// 		queue = append(
// 			queue[1:],
// 			nodeItem{node: node.Left, rank: item.rank + 1},
// 			nodeItem{node: node.Right, rank: item.rank + 1},
// 		)
// 	}
// 	return -1
// }
