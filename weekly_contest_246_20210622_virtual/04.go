// 1906. Minimum Absolute Difference Queries
// https://leetcode.com/contest/weekly-contest-246/problems/minimum-absolute-difference-queries/
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	fmt.Println(minDifference([]int{1, 3, 4, 8}, [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 3}}))
	fmt.Println([]int{2, 1, 4, 1})
	pp.Println("=========================================")
	fmt.Println(minDifference([]int{4, 5, 2, 2, 7, 10}, [][]int{{2, 3}, {0, 2}, {0, 5}, {3, 5}}))
	fmt.Println([]int{-1, 1, 1, 3})
	pp.Println("=========================================")
}

/**
 * すべて同じ数字の場合のみ 0
 * 部分的に同じ数字があっても無視
 */

// 2 <= nums.length <= 105
// 1 <= nums[i] <= 100
// 1 <= queries.length <= 2 * 104
// 0 <= li < ri < nums.length
func minDifference(nums []int, queries [][]int) []int {
	// segTree := make([][2]int, 131072) // 2^7
	return nil
}
