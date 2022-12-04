// 5855. Find the Kth Largest Integer in the Array
// https://leetcode.com/contest/weekly-contest-256/problems/find-the-kth-largest-integer-in-the-array/
package main

import (
	"sort"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(kthLargestNumber([]string{"3", "6", "7", "10"}, 4))
	pp.Println("3")
	pp.Println("=========================================")
	pp.Println(kthLargestNumber([]string{"2", "21", "12", "1"}, 3))
	pp.Println("2")
	pp.Println("=========================================")
	pp.Println(kthLargestNumber([]string{"0", "0"}, 2))
	pp.Println("0")
	pp.Println("=========================================")
}

/**
 * [description]
 * - Time Complexity: O(X)
 * - Space Complexity: O(X)
 */

type numStrings []string

func (a numStrings) Len() int {
	return len(a)
}
func (a numStrings) Less(i, j int) bool {
	return len(a[i]) > len(a[j]) || (len(a[i]) == len(a[j]) && a[i] > a[j])
}
func (a numStrings) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func kthLargestNumber(nums []string, k int) string {
	ns := numStrings(nums)
	sort.Sort(ns)
	return string(ns[k-1])
}
