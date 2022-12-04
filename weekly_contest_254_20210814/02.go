// 5832. Array With Elements Not Equal to Average of Neighbors
// https://leetcode.com/contest/weekly-contest-254/problems/array-with-elements-not-equal-to-average-of-neighbors/
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
}

func rearrangeArray(nums []int) []int {
	sort.Ints(nums)
	ans := make([]int, len(nums))
	left, right := 0, len(nums)-1
	for i := range ans {
		if i%2 == 0 {
			ans[i] = nums[left]
			left++
		} else {
			ans[i] = nums[right]
			right--
		}
	}
	return ans
}
