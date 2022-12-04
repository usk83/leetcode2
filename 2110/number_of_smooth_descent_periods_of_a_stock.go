// 2110. Number of Smooth Descent Periods of a Stock
// https://leetcode.com/problems/number-of-smooth-descent-periods-of-a-stock/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(getDescentPeriods([]int{3, 2, 1, 4}))
	pp.Println(7)
	pp.Println("=========================================")
	pp.Println(getDescentPeriods([]int{8, 6, 7, 7}))
	pp.Println(4)
	pp.Println("=========================================")
	pp.Println(getDescentPeriods([]int{1}))
	pp.Println(1)
	pp.Println("=========================================")
}

/**
 * Constraints:
 * - 1 <= prices.length <= 10^5
 * - 1 <= prices[i] <= 10^5
 *
 * N: Length of the given array `prices`
 */

/**
 * 1. DP
 * - Time Complexity: O(N)
 * - Space Complexity: O(1)
 * > Runtime: 128 ms (faster than 93.75%)
 * > Memory Usage: 9.3 MB (less than 84.38%)
 */
func getDescentPeriods(prices []int) int64 {
	sum, prev := int64(1), 1
	for i := range prices[1:] {
		if prices[i+1]+1 == prices[i] {
			prev++
		} else {
			prev = 1
		}
		sum += int64(prev)
	}
	return sum
}
