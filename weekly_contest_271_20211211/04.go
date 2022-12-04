// 5955. Maximum Fruits Harvested After at Most K Steps
// https://leetcode.com/contest/weekly-contest-271/problems/maximum-fruits-harvested-after-at-most-k-steps/
package main

import (
	"sort"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(maxTotalFruits([][]int{{0, 9}, {4, 1}, {5, 7}, {6, 2}, {7, 4}, {10, 9}}, 5, 40))
	pp.Println(14)
	pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
}

/**
 * Constraints:
 * - 1 <= fruits.length <= 10^5
 * - fruits[i].length == 2
 * - 0 <= startPos, positioni <= 2 * 10^5
 * - positioni-1 < positioni for any i > 0 (0-indexed)
 * - 1 <= amounti <= 10^4
 * - 0 <= k <= 2 * 10^5
 */

func maxTotalFruits(fruits [][]int, startPos int, k int) int {
	var maxTotal int

	rightEnd := sort.Search(len(fruits), func(i int) bool { return fruits[i][0] > startPos+k }) - 1
	leftEnd := sort.Search(len(fruits), func(i int) bool { return fruits[i][0] >= startPos-k })

	left := sort.Search(len(fruits), func(i int) bool { return fruits[i][0] > startPos }) - 1
	var rightTotal int
	rightStart := left + 1
	for i := rightStart; i <= rightEnd; i++ {
		rightTotal += fruits[i][1]
	}
	curTotal := rightTotal
	maxTotal = max(maxTotal, curTotal)
	for left >= 0 && left >= (k-(rightStart-startPos))/2 {
		curTotal += fruits[left][1]
		rightEdge := startPos + k - (startPos-fruits[left][0])*2
		for fruits[rightEnd][0] > rightEdge {
			curTotal -= fruits[rightEnd][1]
			rightEnd--
		}
		left--
		maxTotal = max(maxTotal, curTotal)
	}

	return maxTotal
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
