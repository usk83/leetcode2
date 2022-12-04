// 26. Remove Duplicates from Sorted Array
// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
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
 * - 0 <= nums.length <= 3 * 10^4
 * - -100 <= nums[i] <= 100
 * - nums is sorted in non-decreasing order.
 *
 * N: Length of the input array `nums`
 */

/**
 * 1-1. Generalized solution without magic numbers
 */
// func removeDuplicates(nums []int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}
// 	p := 1
// 	for _, n := range nums[1:] {
// 		if n != nums[p-1] {
// 			nums[p] = n
// 			p++
// 		}
// 	}
// 	return p
// }

/**
 * 1. The original solution
 * - Time Complexity: O(N)
 * - Space Complexity: O(1)
 * > Runtime: 4 ms (faster than 99.30%)
 * > Memory Usage: 4.4 MB (less than 42.82%)
 */
func removeDuplicates(nums []int) int {
	p, prev := 0, -101 // The minimum possible value in nums is -100
	for _, n := range nums {
		if n != prev {
			nums[p] = n
			prev = n
			p++
		}
	}
	return p
}
