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
 * [description]
 * - Time Complexity: O(X)
 * - Space Complexity: O(X)
 * > Runtime: X ms (faster than 100.00%)
 * > Memory Usage: X MB (less than 100.00%)
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	return -1
}
