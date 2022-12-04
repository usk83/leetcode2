// 215. Kth Largest Element in an Array
// https://leetcode.com/problems/kth-largest-element-in-an-array/
package main

import (
	"math/rand"
	"time"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	pp.Println(5)
	pp.Println("=========================================")
	pp.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
	pp.Println(4)
	pp.Println("=========================================")
	pp.Println(findKthLargest([]int{5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 3, 2, 1}, 8))
	pp.Println(5)
	pp.Println("=========================================")
	pp.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	pp.Println(5)
	pp.Println("=========================================")
}

/**
 * Constraints:
 * - 1 <= k <= nums.length <= 10^4
 * - -10^4 <= nums[i] <= 10^4
 *
 * K: Given value to determine the element to return
 * N: Length of the input array `nums`
 */

/**
 * 3. Shuffle the array, then Quick Select
 * - Time Complexity: O(N)
 * - Space Complexity: O(1)
 * > Runtime: 4 ms (faster than 98.32%)
 * > Memory Usage: 3.5 MB (less than 100.00%)
 */
func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(nums), func(i, j int) { nums[i], nums[j] = nums[j], nums[i] })
	for len(nums) > 1 {
		left := 0
		for pivot, right := min(nums[0], nums[1]), len(nums)-1; left < right; {
			if nums[left] > pivot {
				left++
			} else if nums[right] <= pivot {
				right--
			} else {
				nums[left], nums[right] = nums[right], nums[left]
			}
		}
		if left == 0 {
			if k > 2 {
				left = 2
			} else {
				left = 1
			}
		}
		if k <= left {
			nums = nums[:left]
		} else {
			nums = nums[left:]
			k -= left
		}
	}
	return nums[0]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
