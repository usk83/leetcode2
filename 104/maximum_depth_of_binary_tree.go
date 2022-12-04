// 104. Maximum Depth of Binary Tree
// https://leetcode.com/problems/maximum-depth-of-binary-tree/
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
 * [description]
 * - Time Complexity: O(N)
 * - Space Complexity: O(N)
 * > Runtime: X ms (faster than 100.00%)
 * > Memory Usage: X MB (less than 100.00%)
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func maxDepth(node *TreeNode) int {
// 	if node == nil {
// 		return 0
// 	}
// 	return max(maxDepth(node.Left), maxDepth(node.Right)) + 1
// }

// func maxDepth(root *TreeNode) int {
// 	var dfs func(*TreeNode) int
// 	dfs = func(node *TreeNode) int {
// 		if node == nil {
// 			return 0
// 		}
// 		return max(dfs(node.Left), dfs(node.Right)) + 1
// 	}
// 	return dfs(root)
// }

// func maxDepth(root *TreeNode) int {
// 	var deepest int
// 	var dfs func(*TreeNode, int)
// 	dfs = func(node *TreeNode, level int) {
// 		if node == nil {
// 			return
// 		}
// 		level += 1
// 		if node.Left == nil && node.Right == nil {
// 			deepest = max(deepest, level)
// 		} else {
// 			dfs(node.Left, level)
// 			dfs(node.Right, level)
// 		}
// 	}
// 	dfs(root, 0)
// 	return deepest
// }

func maxDepth(root *TreeNode) int {
	var dfs func(*TreeNode, int) int
	dfs = func(node *TreeNode, level int) int {
		if node == nil {
			return level
		}
		level += 1
		return max(dfs(node.Left, level), dfs(node.Right, level))
	}
	return dfs(root, 0)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
