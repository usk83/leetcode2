// 200. Number of Islands
// https://leetcode.com/problems/number-of-islands/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(numIslands([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(numIslands([][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}))
	pp.Println(3)
	pp.Println("=========================================")
}

/**
 * Constraints:
 * - m == grid.length
 * - n == grid[i].length
 * - 1 <= m, n <= 300
 * - grid[i][j] is '0' or '1'.
 *
 * N: Width of a given `grid`
 * M: Height of a given `grid`
 */

const (
	water = '0'
	land  = '1'
)

var directions = [4][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

/**
 * 2. BFS
 * - Time Complexity: O(NM)
 * - Space Complexity: O(NM)
 * > Runtime: 4 ms (faster than 47.83%)
 * > Memory Usage: 5.5 MB (less than 17.64%)
 */
// func numIslands(grid [][]byte) int {
// 	n, m := len(grid[0]), len(grid)
// 	var (
// 		count int
// 		queue [][2]int
// 	)
// 	for y := range grid {
// 		for x := range grid[0] {
// 			if grid[y][x] == water {
// 				continue
// 			}
// 			grid[y][x] = water
// 			queue = append(queue, [2]int{x, y}) // enqueue
// 			for len(queue) != 0 {
// 				col, row := queue[0][0], queue[0][1] // peek
// 				queue = queue[1:]                    // remove (dequeue)
// 				for _, dir := range directions {
// 					if c, r := col+dir[0], row+dir[1]; c >= 0 && c < n && r >= 0 && r < m && grid[r][c] == land {
// 						grid[r][c] = water
// 						queue = append(queue, [2]int{c, r})
// 					}
// 				}
// 			}
// 			count++
// 		}
// 	}
// 	return count
// }

/**
 * 1. DFS
 * - Time Complexity: O(NM)
 * - Space Complexity: O(NM)
 * > Runtime: 0 ms (faster than 100.00%)
 * > Memory Usage: 4.2 MB (less than 24.51%) // efficient enough
 */
func numIslands(grid [][]byte) int {
	n, m := len(grid[0]), len(grid)
	var dfs func(int, int)
	dfs = func(col, row int) {
		grid[row][col] = water
		for _, dir := range directions {
			if c, r := col+dir[0], row+dir[1]; c >= 0 && c < n && r >= 0 && r < m && grid[r][c] == land {
				dfs(c, r)
			}
		}
	}
	var count int
	for row := range grid {
		for col := range grid[0] {
			if grid[row][col] == land {
				dfs(col, row)
				count++
			}
		}
	}
	return count
}
