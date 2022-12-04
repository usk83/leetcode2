// 5854. Minimum Difference Between Highest and Lowest of K Scores
// https://leetcode.com/contest/weekly-contest-256/problems/minimum-difference-between-highest-and-lowest-of-k-scores/
package main

import (
	"sort"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(minimumDifference([]int{90}, 1))
	pp.Println(0)
	pp.Println("=========================================")
	pp.Println(minimumDifference([]int{9, 4, 1, 7}, 2))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(minimumDifference([]int{9, 15, 1, 7, 33, 53, 54, 55}, 3))
	pp.Println()
	pp.Println("=========================================")

}

/**
 * [description]
 * - Time Complexity: O(X)
 * - Space Complexity: O(X)
 */
func minimumDifference(nums []int, k int) int {
	if k < 2 {
		return 0
	}
	sort.Ints(nums)
	diff := nums[k-1] - nums[0]
	for i := len(nums) - 1; i > 0; i-- {
		nums[i] = nums[i] - nums[i-1]
	}
	minDiff := diff
	for i := k; i < len(nums); i++ {
		diff += nums[i] - nums[i-k+1]
		minDiff = min(minDiff, diff)
	}
	return minDiff
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
