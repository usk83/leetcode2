// 46. Permutations
// https://leetcode.com/problems/permutations/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(permute([]int{1, 2, 3}))
	pp.Println([][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}})
	pp.Println("=========================================")
	pp.Println(permute([]int{0, 1}))
	pp.Println([][]int{{0, 1}, {1, 0}})
	pp.Println("=========================================")
	pp.Println(permute([]int{1}))
	pp.Println([][]int{{1}})
	pp.Println("=========================================")
}

/**
 * Constraints:
 * - 1 <= nums.length <= 6
 * - -10 <= nums[i] <= 10
 * - All the integers of nums are unique
 *
 * N: Length of the input array `nums`
 */

/**
 * 1. Backtracking
 * - Time Complexity: O(N!)
 * - Space Complexity: O(N!)
 * > Runtime: 0 ms (faster than 100.00%)
 * > Memory Usage: 2.6 MB (less than 86.90%)
 */
func permute(nums []int) [][]int {
	n, permutations := len(nums), [][]int{}
	var bk func([]int, []bool)
	bk = func(perm []int, used []bool) {
		if len(perm) == n {
			p := make([]int, n)
			copy(p, perm)
			permutations = append(permutations, p)
			return
		}
		for i := range nums {
			if used[i] {
				continue
			}
			used[i] = true
			perm = append(perm, nums[i])
			bk(perm, used)
			perm = perm[:len(perm)-1]
			used[i] = false
		}
	}
	bk(make([]int, 0, n), make([]bool, n))
	return permutations
}
