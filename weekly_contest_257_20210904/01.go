// 5863. Count Special Quadruplets
// https://leetcode.com/contest/weekly-contest-257/problems/count-special-quadruplets/
package main

import (
	"sort"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(countQuadruplets([]int{1, 2, 3, 6}))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(countQuadruplets([]int{3, 3, 6, 4, 5}))
	pp.Println(0)
	pp.Println("=========================================")
	pp.Println(countQuadruplets([]int{1, 1, 1, 3, 5}))
	pp.Println(4)
	pp.Println("=========================================")
	pp.Println(countQuadruplets([]int{9, 6, 8, 23, 39, 23}))
	pp.Println(2)
	pp.Println("=========================================")
}

/**
 * Constraints:
 * - 4 <= nums.length <= 50
 * - 1 <= nums[i] <= 100
 */

/**
 * [description]
 * - Time Complexity: O(X)
 * - Space Complexity: O(X)
 * > Runtime: X ms (faster than X%)
 * > Memory Usage: X MB (less than X%)
 */

/**
 * 2.
 * - Time Complexity: O(N^3 * logN)
 * - Space Complexity: O(N)
 */
func countQuadruplets(nums []int) int {
	m := map[int][]int{}
	for i, num := range nums {
		m[num] = append(m[num], i)
	}
	var count int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				sum := nums[i] + nums[j] + nums[k]
				if indices, ok := m[sum]; ok {
					count += len(indices) - sort.Search(len(indices), func(i int) bool { return indices[i] > k })
				}
			}
		}
	}
	return count
}

/**
 * 1-3. Refactored B
 * - Time Complexity: O(N^4)
 * - Space Complexity: O(1)
 */
// func countQuadruplets(nums []int) int {
// 	var count int
// 	for i, num1 := range nums[:len(nums)-3] {
// 		for j, num2 := range nums[i+1 : len(nums)-2] {
// 			for k, num3 := range nums[i+j+2 : len(nums)-1] {
// 				for _, num4 := range nums[i+j+k+3:] {
// 					if num1+num2+num3 == num4 {
// 						count++
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return count
// }

/**
 * 1-2. Refactored A
 * - Time Complexity: O(N^4)
 * - Space Complexity: O(1)
 */
// func countQuadruplets(nums []int) int {
// 	var count int
// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			for k := j + 1; k < len(nums); k++ {
// 				for l := k + 1; l < len(nums); l++ {
// 					if nums[i]+nums[j]+nums[k] == nums[l] {
// 						count++
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return count
// }

/**
 * 1. Original solution
 * - Time Complexity: O(N^4)
 * - Space Complexity: O(1)
 * > Runtime: 8 ms (faster than 95.83%)
 * > Memory Usage: 2.3 MB (less than 100%)
 */
// func countQuadruplets(nums []int) int {
// 	var count int
// 	for i := 0; i < len(nums); i++ {
// 		sum := nums[i]
// 		for j := i + 1; j < len(nums); j++ {
// 			sum += nums[j]
// 			for k := j + 1; k < len(nums); k++ {
// 				sum += nums[k]
// 				for l := k + 1; l < len(nums); l++ {
// 					if sum == nums[l] {
// 						count++
// 					}
// 				}
// 				sum -= nums[k]
// 			}
// 			sum -= nums[j]
// 		}
// 		sum -= nums[i]
// 	}
// 	return count
// }
