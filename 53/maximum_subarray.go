// 53. Maximum Subarray
// https://leetcode.com/problems/maximum-subarray/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	pp.Println(6)
	pp.Println("=========================================")
	pp.Println(maxSubArray([]int{1}))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(maxSubArray([]int{5, 4, -1, 7, 8}))
	pp.Println(23)
	pp.Println("=========================================")
	pp.Println(maxSubArray([]int{-2, 1}))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(maxSubArray([]int{-2, -3, -1}))
	pp.Println(-1)
	pp.Println("=========================================")
}

/**
 * Constraints:
 * - 1 <= nums.length <= 10^5
 * - -10^4 <= nums[i] <= 10^4
 */

/**
 * 3. Divide and conquer
 * - Time Complexity: O(N)?
 * - Space Complexity: O(logN)?
 * > Runtime: 124 ms (faster than 28.63%)
 * > Memory Usage: 10.4 MB (less than 5.85%)
 */
// type subarray struct {
// 	sum      int
// 	leftSum  int
// 	rightSum int
// 	maxSum   int
// }
//
// func maxSubArray(nums []int) int {
// 	var dc func(int, int) *subarray
// 	dc = func(l, r int) *subarray {
// 		if l+1 == r {
// 			return &subarray{nums[l], nums[l], nums[l], nums[l]}
// 		}
// 		m := (l + r) / 2
// 		left, right := dc(l, m), dc(m, r)
// 		return &subarray{
// 			left.sum + right.sum,
// 			max(left.leftSum, left.sum+right.leftSum),
// 			max(right.rightSum, right.sum+left.rightSum),
// 			max(max(left.maxSum, right.maxSum), left.rightSum+right.leftSum),
// 		}
// 	}
// 	return dc(0, len(nums)).maxSum
// }

/**
 * 2-2. Optimized DP
 * - Time Complexity: O(N)
 * - Space Complexity: O(1)
 * > Runtime: 96 ms (faster than 91.59%)
 * > Memory Usage: 9.5 MB (less than 28.47%)
 */
func maxSubArray(nums []int) int {
	maxSum, curSum := nums[0], nums[0]
	for _, num := range nums[1:] {
		curSum = max(curSum+num, num)
		maxSum = max(maxSum, curSum)
	}
	return maxSum
}

/**
 * 2-1. Straight forward DP
 * - Time Complexity: O(N)
 * - Space Complexity: O(N)
 */
// func maxSubArray(nums []int) int {
// 	dp := make([]int, len(nums))
// 	dp[0] = nums[0]
// 	for i, num := range nums[1:] {
// 		dp[i+1] = max(dp[i]+num, num)
// 	}
// 	maxSum := dp[0]
// 	for _, sum := range dp[1:] {
// 		maxSum = max(maxSum, sum)
// 	}
// 	return maxSum
// }

/**
 * 1. Greedy
 * - Time Complexity: O(N)
 * - Space Complexity: O(1)
 */
// func maxSubArray(nums []int) int {
// 	maxSum, curSum := nums[0], nums[0]
// 	for _, num := range nums[1:] {
// 		curSum = max(curSum, 0)
// 		curSum += num
// 		maxSum = max(maxSum, curSum)
// 	}
// 	return maxSum
// }

/**
 * 0-2. Bit improved
 * - Time Complexity: O(N)
 * - Space Complexity: O(1)
 */
// func maxSubArray(nums []int) int {
// 	maxSum, curSum := MinInt, 0
// 	for _, num := range nums {
// 		curSum += num
// 		if curSum > maxSum {
// 			maxSum = curSum
// 		}
// 		if curSum < 0 {
// 			curSum = 0
// 		}
// 	}
// 	return maxSum
// }

/**
 * 0-1. First solution
 * - Time Complexity: O(N)
 * - Space Complexity: O(1)
 */
// func maxSubArray(nums []int) int {
// 	maxNum := MinInt
// 	for _, num := range nums {
// 		maxNum = max(maxNum, num)
// 	}
// 	if maxNum <= 0 {
// 		return maxNum
// 	}
//
// 	var maxSum, curSum int
// 	for _, num := range nums {
// 		curSum += num
// 		if curSum < 0 {
// 			curSum = 0
// 		} else {
// 			maxSum = max(maxSum, curSum)
// 		}
// 	}
// 	return maxSum
// }

// ========================
// ========================
// ========================

const (
	MaxInt = int(^uint(0) >> 1)
	MinInt = -MaxInt - 1
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
