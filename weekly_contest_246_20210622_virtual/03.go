// 1905. Count Sub Islands
// https://leetcode.com/contest/weekly-contest-246/problems/count-sub-islands/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(countSubIslands(
		[][]int{{1, 1, 1, 0, 0}, {0, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 1, 1}},
		[][]int{{1, 1, 1, 0, 0}, {0, 0, 1, 1, 1}, {0, 1, 0, 0, 0}, {1, 0, 1, 1, 0}, {0, 1, 0, 1, 0}},
	))
	pp.Println(3)
	pp.Println("=========================================")
	pp.Println(countSubIslands(
		[][]int{{1, 0, 1, 0, 1}, {1, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 1, 1, 1, 1}, {1, 0, 1, 0, 1}},
		[][]int{{0, 0, 0, 0, 0}, {1, 1, 1, 1, 1}, {0, 1, 0, 1, 0}, {0, 1, 0, 1, 0}, {1, 0, 0, 0, 1}},
	))
	pp.Println(2)
	pp.Println("=========================================")
}

/**
 * grid2の方でdfsしながら島を数える
 *   海に変えていく
 *   全部変えきる
 * 途中で無効になったかを確認しておく
 */

// m == grid1.length == grid2.length
// n == grid1[i].length == grid2[i].length
// 1 <= m, n <= 500
// grid1[i][j] and grid2[i][j] are either 0 or 1.
func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	count := 0

	var dfs func(v, h int) bool
	dfs = func(v, h int) bool {
		if v < 0 || v >= len(grid2) || h < 0 || h >= len(grid2[0]) || grid2[v][h] == 0 {
			return true
		}

		grid2[v][h] = 0

		ok := true
		if grid1[v][h] != 1 {
			ok = false
		}

		ok = dfs(v-1, h) && ok
		ok = dfs(v, h+1) && ok
		ok = dfs(v+1, h) && ok
		ok = dfs(v, h-1) && ok

		return ok
	}

	for v := 0; v < len(grid2); v++ {
		for h := 0; h < len(grid2[0]); h++ {
			if grid2[v][h] == 1 {
				if dfs(v, h) {
					count++
				}
			}
		}
	}

	return count
}
