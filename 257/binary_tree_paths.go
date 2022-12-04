// 257. Binary Tree Paths
// https://leetcode.com/problems/binary-tree-paths/
package main

import (
	"strconv"
	"strings"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println("No local tests available.")
	pp.Println("=========================================")

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * v3
 */
func binaryTreePaths(root *TreeNode) []string {
	paths := []string{}

	var dfs func(node *TreeNode, pathVals []string) []string
	dfs = func(node *TreeNode, pathVals []string) []string {
		if node == nil {
			return
		}
		pathVals = append(pathVals, strconv.Itoa(node.Val))
		if node.Left == nil && node.Right == nil {
			paths = append(paths, strings.Join(pathVals, "->"))
		} else {
			pathVals = dfs(node.Left, pathVals)
			pathVals = dfs(node.Right, pathVals)
		}
		pathVals = pathVals[:len(pathVals)-1]
		return pathVals
	}
	_ = dfs(root, []string{})

	return paths
}

/**
 * v2. Keep the path before as a slice of string
 */
// func binaryTreePaths(root *TreeNode) []string {
// 	paths := []string{}
//
// 	var dfs func(node *TreeNode, pathVals []string)
// 	dfs = func(node *TreeNode, pathVals []string) {
// 		if node == nil {
// 			return
// 		}
// 		pathVals = append(pathVals, strconv.Itoa(node.Val))
// 		if node.Left == nil && node.Right == nil {
// 			paths = append(paths, strings.Join(pathVals, "->"))
// 			return
// 		}
// 		dfs(node.Left, pathVals)
// 		dfs(node.Right, pathVals)
//
// 	}
// 	dfs(root, []string{})
//
// 	return paths
// }

/**
 * v1. Keep the path before as one string
 */
// func binaryTreePaths(root *TreeNode) []string {
// 	paths := []string{}
//
// 	var dfs func(node *TreeNode, path string)
// 	dfs = func(node *TreeNode, path string) {
// 		// 存在しないノードの処理して
// 		if node == nil {
// 			return
// 		}
//
// 		// ここまでの文字列を作って
// 		path += "->" + strconv.Itoa(node.Val)
//
// 		// 末尾だったら`paths`に追加して終了
// 		if node.Left == nil && node.Right == nil {
// 			paths = append(paths, strings.TrimPrefix(path, "->"))
// 			return
// 		}
//
// 		// そうじゃなかったら左右呼び出し
// 		dfs(node.Left, path)
// 		dfs(node.Right, path)
// 	}
//
// 	dfs(root, "")
//
// 	return paths
// }
