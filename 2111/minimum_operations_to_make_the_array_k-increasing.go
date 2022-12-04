// 2111. Minimum Operations to Make the Array K-Increasing
// https://leetcode.com/problems/minimum-operations-to-make-the-array-k-increasing/
package main

import (
	"sort"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(kIncreasing([]int{5, 4, 3, 2, 1}, 1))
	pp.Println(4)
	pp.Println("=========================================")
	pp.Println(kIncreasing([]int{4, 1, 5, 2, 6, 2}, 2))
	pp.Println(0)
	pp.Println("=========================================")
	pp.Println(kIncreasing([]int{4, 1, 5, 2, 6, 2}, 3))
	pp.Println(2)
	pp.Println("=========================================")
}

/**
 * Constraints:
 * - 1 <= arr.length <= 10^5
 * - 1 <= arr[i], k <= arr.length
 *
 * N: Length of the given array `arr`
 */

/**
 * 1. Longest Non-Decreasing Subsequence
 * - Time Complexity: O((N/K)logN)
 * - Space Complexity: O(N/K)
 * > Runtime: 156 ms (faster than 85.71%)
 * > Memory Usage: 10 MB (less than 92.86%)
 */
func kIncreasing(arr []int, k int) int {
	op := len(arr)
	for i := 1; i <= k; i++ {
		dp := []int{}
		for j := len(arr) - i; j >= 0; j = j - k {
			idx := sort.Search(len(dp), func(i int) bool {
				return arr[j] > dp[i]
			})
			if idx == len(dp) {
				dp = append(dp, arr[j])
			} else {
				dp[idx] = arr[j]
			}
		}
		op -= len(dp)
	}
	return op
}
