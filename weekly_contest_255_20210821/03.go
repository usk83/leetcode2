// 5852. Minimize the Difference Between Target and Chosen Elements
// https://leetcode.com/contest/weekly-contest-255/problems/minimize-the-difference-between-target-and-chosen-elements/
package main

import (
	"sort"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(minimizeTheDifference(
		[][]int{
			{7, 8, 4},
			{2, 9, 6},
			{1, 5, 3},
		},
		13,
	))
	pp.Println(0)
	pp.Println("=========================================")
	pp.Println(minimizeTheDifference([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 13))
	pp.Println(0)
	pp.Println("=========================================")
	pp.Println(minimizeTheDifference([][]int{{1}, {2}, {3}}, 100))
	pp.Println(94)
	pp.Println("=========================================")
	pp.Println(minimizeTheDifference([][]int{{1, 2, 9, 8, 7}}, 6))
	pp.Println(1)
	pp.Println("=========================================")
}

/**
 * [description]
 * - Time Complexity: O(X)
 * - Space Complexity: O(X)
 */
// m == mat.length, n == mat[i].length
// 1 <= m, n <= 70
// 1 <= mat[i][j] <= 70
// 1 <= target <= 800
func minimizeTheDifference(mat [][]int, target int) int {
	m, n := len(mat), len(mat[0])
	_, _ = m, n
	var smallest int
	for row := range mat {
		sort.Ints(mat[row])
		min := mat[row][0]
		smallest += min
		for col := range mat[row] {
			mat[row][col] -= min
		}
	}

	if smallest >= target {
		return smallest - target
	}

	ideal := target - smallest
	// // dpLen := min(70, ideal)
	dpLen := (ideal) + (ideal)>>1 + 1
	dp := make([]bool, dpLen)
	dp[0] = true
	// fmt.Println(-1, dp)
	for row := len(mat) - 1; row >= 0; row-- {
		newDP := make([]bool, dpLen)
		for org := range dp {
			if !dp[org] {
				continue
			}
			for _, add := range mat[row] {
				next := org + add
				if next < dpLen {
					newDP[next] = true
				}
			}
		}
		dp = newDP
		// fmt.Println(row, dp)
	}

	best := 0
	for num := range dp {
		if !dp[num] {
			continue
		}
		if num < ideal {
			best = num
			continue
		}
		if num-ideal < ideal-best {
			best = num
		}
		break
	}

	// pp.Println("=== DEBUG START ======================================")
	// pp.Println("target:", target)
	// pp.Println("smallest:", smallest)
	// pp.Println("ideal:", ideal)
	// pp.Println("best:", best)
	// fmt.Println(mat)
	// fmt.Println(dp)
	// pp.Println("=== DEBUG END ========================================")

	if best < ideal {
		return ideal - best
	}
	return best - ideal
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
