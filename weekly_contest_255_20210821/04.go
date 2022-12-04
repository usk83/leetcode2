// 5853. Find Array Given Subset Sums
// https://leetcode.com/contest/weekly-contest-255/problems/find-array-given-subset-sums/
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	fmt.Println(recoverArray(3, []int{-3, -2, -1, 0, 0, 1, 2, 3}))
	fmt.Println([]int{1, 2, -3})
	pp.Println("=========================================")
	fmt.Println(recoverArray(2, []int{0, 0, 0, 0}))
	fmt.Println([]int{0, 0})
	pp.Println("=========================================")
	fmt.Println(recoverArray(4, []int{0, 0, 5, 5, 4, -1, 4, 9, 9, -1, 4, 3, 4, 8, 3, 8}))
	fmt.Println([]int{0, -1, 4, 5})
	pp.Println("=========================================")
}

/**
 * n = 4
 * sums = [0,0,5,5,4,-1,4,9,9,-1,4,3,4,8,3,8]
 *
 * [-1, -1, 0, 0, 3, 3, 4, 4, 4, 4, 5, 5, 8, 8, 9, 9]
 *
 * -1
 * 0
 * 4
 * 5
 *
 *
 * Output: [0,-1,4,5]
 * Explanation: [0,-1,4,5] is able to achieve the given subset sums.
 *
 */

/**
 * [description]
 * - Time Complexity: O(X)
 * - Space Complexity: O(X)
 */
// 1 <= n <= 15
// sums.length == 2^n == 32768
// -10^4 <= sums[i] <= 10^4
func recoverArray(n int, sums []int) []int {
	return nil
}
