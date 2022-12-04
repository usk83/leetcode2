// 35. Search Insert Position
// https://leetcode.com/problems/search-insert-position/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(searchInsert([]int{1, 3, 5, 6}, 7))
	pp.Println(4)
	pp.Println("=========================================")
	pp.Println(searchInsert([]int{1, 3, 5, 6}, 0))
	pp.Println(0)
	pp.Println("=========================================")
	pp.Println(searchInsert([]int{1}, 0))
	pp.Println(0)
	pp.Println("=========================================")
}

/**
 * Constraints:
 * - 1 <= nums.length <= 10^4
 * - -10^4 <= nums[i] <= 10^4
 * - nums contains distinct values sorted in ascending order.
 * - -10^4 <= target <= 10^4
 */

/**
 * 2. Binary Search
 * - Time Complexity: O(logN)
 * - Space Complexity: O(1)
 * > Runtime: 4 ms (faster than 79.78%)
 * > Memory Usage: 3 MB (less than 42.52%)
 */
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		if center := left + (right-left)>>1; nums[center] < target {
			left = center + 1
		} else {
			right = center
		}
	}
	return right
}

/**
 * 1. Use `sort.Search`
 * - Time Complexity: O(logN)
 * - Space Complexity: O(1)
 * > Runtime: 0 ms (faster than 100.00%)
 * > Memory Usage: 3 MB (less than 42.52%)
 */
// func searchInsert(nums []int, target int) int {
// 	return sort.Search(len(nums), func(i int) bool { return target <= nums[i] })
// }
