// 695. Max Area of Island
// https://leetcode.com/problems/max-area-of-island/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(maxAreaOfIsland([][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}))
	pp.Println(6)
	pp.Println("=========================================")
	pp.Println(maxAreaOfIsland([][]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
	}))
	pp.Println(0)
	pp.Println("=========================================")
}

/**
 * Constraints:
 * - m == grid.length
 * - n == grid[i].length
 * - 1 <= m, n <= 50
 * - grid[i][j] is either 0 or 1.
 */

/**
 * 2. Check conditions earlier to reduce function calls
 * - Time Complexity: O(n*m)
 * - Space Complexity: O(n*m)
 * > Runtime: 8 ms (faster than 98.01%)
 * > Memory Usage: 5.1 MB (less than 64.38%)
 */
func maxAreaOfIsland(grid [][]int) int {
	n, m := len(grid[0]), len(grid)

	var dfs func(int, int) int
	dfs = func(x, y int) int {
		grid[y][x] = 0
		count := 1
		if y > 0 && grid[y-1][x] == 1 {
			count += dfs(x, y-1)
		}
		if x < n-1 && grid[y][x+1] == 1 {
			count += dfs(x+1, y)
		}
		if y < m-1 && grid[y+1][x] == 1 {
			count += dfs(x, y+1)
		}
		if x > 0 && grid[y][x-1] == 1 {
			count += dfs(x-1, y)
		}
		return count
	}

	var maxArea int
	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			if grid[y][x] == 1 {
				if area := dfs(x, y); area > maxArea {
					maxArea = area
				}
			}
		}
	}
	return maxArea
}

/**
 * 1. Count the area of an island while submerging the land in water
 * - Time Complexity: O(n*m)
 * - Space Complexity: O(n*m)
 * > Runtime: 12 ms (faster than 79.42%)
 * > Memory Usage: 5 MB (less than 77.65%)
 */
// func maxAreaOfIsland(grid [][]int) int {
// 	n, m := len(grid[0]), len(grid)
//
// 	var dfs func(int, int) int
// 	dfs = func(x, y int) int {
// 		if x < 0 || x >= n || y < 0 || y >= m || grid[y][x] == 0 {
// 			return 0
// 		}
// 		grid[y][x] = 0
// 		return 1 + dfs(x, y-1) + dfs(x+1, y) + dfs(x, y+1) + dfs(x-1, y)
// 	}
//
// 	var maxArea int
// 	for y := 0; y < m; y++ {
// 		for x := 0; x < n; x++ {
// 			maxArea = max(maxArea, dfs(x, y))
// 		}
// 	}
// 	return maxArea
// }
//
// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
