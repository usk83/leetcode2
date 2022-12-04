// 5868?. Number of Pairs of Interchangeable Rectangles
// https://leetcode.com/contest/weekly-contest-258/problems/number-of-pairs-of-interchangeable-rectangles/
package main

import (
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
}

/**
 * [description]
 * - Time Complexity: O(X)
 * - Space Complexity: O(X)
 * > Runtime: X ms (faster than X%)
 * > Memory Usage: X MB (less than X%)
 */
func interchangeableRectangles(rectangles [][]int) int64 {
	cache := map[float64]int{}
	ret := 0
	for i := len(rectangles) - 1; i >= 0; i-- {
		ratio := float64(rectangles[i][0]) / float64(rectangles[i][1])
		if count, ok := cache[ratio]; ok {
			ret += count
			cache[ratio]++
		} else {
			cache[ratio] = 1
		}
	}
	return int64(ret)
}
