// 5800. Build Array from Permutation
// https://leetcode.com/contest/weekly-contest-248/problems/build-array-from-permutation/
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	fmt.Println(buildArray([]int{0, 2, 1, 5, 3, 4}))
	fmt.Println([]int{0, 1, 2, 4, 5, 3})
	pp.Println("=========================================")
	fmt.Println(buildArray([]int{5, 0, 1, 2, 3, 4}))
	fmt.Println([]int{4, 5, 0, 1, 2, 3})
	pp.Println("=========================================")
}

func buildArray(nums []int) []int {
	ans := make([]int, len(nums))
	for i, num := range nums {
		ans[i] = nums[num]
	}
	return ans
}
