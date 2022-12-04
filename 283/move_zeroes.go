// 283. Move Zeroes
// https://leetcode.com/problems/move-zeroes/
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
	fmt.Println([]int{1, 3, 12, 0, 0})
	pp.Println("=========================================")
	nums = []int{0}
	moveZeroes(nums)
	fmt.Println(nums)
	fmt.Println([]int{0})
	pp.Println("=========================================")
}

/**
 * Constraints:
 * - 1 <= nums.length <= 10^4
 * - -2^31 <= nums[i] <= 2^31 - 1
 *
 * N: Length of the input array `nums`
 */

/**
 * 2-2. コードを短くしたバージョン
 * - Time Complexity: O(N)
 * - Space Complexity: O(1)
 */
// func moveZeroes(nums []int) {
// 	var headIndex int
// 	for i := range nums {
// 		if nums[i] != 0 {
// 			nums[headIndex], nums[i] = nums[i], nums[headIndex]
// 			headIndex++
// 		}
// 	}
// }

/**
 * 2-1. 普通にSwapする
 * - Time Complexity: O(N)
 * - Space Complexity: O(1)
 * > Runtime: 20 ms (faster than 70.76%)
 * > Memory Usage: 6.8 MB (less than 66.40%)
 */
// func moveZeroes(nums []int) {
// 	var headIndex int
// 	for i := range nums {
// 		if nums[i] == 0 {
// 			continue
// 		}
// 		if i != headIndex {
// 			nums[headIndex], nums[i] = nums[i], nums[headIndex]
// 		}
// 		headIndex++
// 	}
// }

/**
 * 1. 前を確定させたあと後ろを0にする
 * - Time Complexity: O(N)
 * - Space Complexity: O(1)
 * > Runtime: 20 ms (faster than 70.76%)
 * > Memory Usage: 6.8 MB (less than 66.40%)
 */
func moveZeroes(nums []int) {
	var headIndex int
	for i := range nums {
		if nums[i] == 0 {
			continue
		}
		if i != headIndex {
			nums[headIndex] = nums[i]
		}
		headIndex++
	}
	for i := headIndex; i < len(nums); i++ {
		nums[i] = 0
	}
}
