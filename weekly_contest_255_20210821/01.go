// 5850. Find Greatest Common Divisor of Array
// https://leetcode.com/contest/weekly-contest-255/problems/find-greatest-common-divisor-of-array/
package main

import (
	"sort"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println("No local tests available.")
	pp.Println("=========================================")
}

/**
 * [description]
 * - Time Complexity: O(X)
 * - Space Complexity: O(X)
 */
func findGCD(nums []int) int {
	sort.Ints(nums)
	smallest, largest := nums[0], nums[len(nums)-1]
	for i := smallest; i > 1; i-- {
		if smallest%i == 0 && largest%i == 0 {
			return i
		}
	}
	return 1
}
