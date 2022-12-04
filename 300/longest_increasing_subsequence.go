// 300. Longest Increasing Subsequence
// https://leetcode.com/problems/longest-increasing-subsequence/
package main

import (
	"sort"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	pp.Println(4)
	pp.Println("=========================================")
	pp.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
	pp.Println(4)
	pp.Println("=========================================")
	pp.Println(lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6}))
	pp.Println(6)
	pp.Println("=========================================")
}

/**
 * [description]
 * - Time Complexity: O(X)
 * - Space Complexity: O(X)
 * > Runtime: X ms (faster than 100.00%)
 * > Memory Usage: X MB (less than 100.00%)
 */

// const (
// 	MaxInt = int(^uint(0) >> 1)
// )

// BinarySearchで検索する

/**
 * 300. Longest Increasing Subsequence
 *
 * 2. DP + BinarySearch
 *
 * N: Length of the given array
 *
 * - Time Complexity: O(NlogN)
 * - Space Complexity: O(N)
 */
const (
	MaxInt = int(^uint(0) >> 1)
)

func lengthOfLIS(nums []int) int {
	dp := []int{MaxInt}
	for i := len(nums) - 1; i >= 0; i-- {
		target := sort.Search(len(dp), func(k int) bool { return dp[k] <= nums[i] })
		if target == len(dp) {
			dp = append(dp, nums[i])
		} else {
			dp[target] = max(nums[i], dp[target])
		}
	}
	return len(dp) - 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// とりあえずDP
// func lengthOfLIS(nums []int) int {
// 	dp := []int{MaxInt}
// 	for i := len(nums) - 1; i >= 0; i-- {
// 		for j := len(dp) - 1; j >= 0; j-- {
// 			if nums[i] < dp[j] {
// 				if j+1 == len(dp) {
// 					dp = append(dp, nums[i])
// 				} else {
// 					dp[j+1] = max(nums[i], dp[j+1])
// 				}
// 				break
// 			}
// 		}
// 	}
// 	return len(dp) - 1
// }

// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }
