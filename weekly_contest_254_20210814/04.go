// 5845. Last Day Where You Can Still Cross
// https://leetcode.com/contest/weekly-contest-254/problems/last-day-where-you-can-still-cross/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(latestDayToCross(2, 2, [][]int{{1, 1}, {2, 1}, {1, 2}, {2, 2}}))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(latestDayToCross(2, 2, [][]int{{1, 1}, {1, 2}, {2, 1}, {2, 2}}))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(latestDayToCross(3, 3, [][]int{{1, 2}, {2, 1}, {3, 3}, {2, 2}, {1, 1}, {1, 3}, {2, 3}, {3, 2}, {3, 1}}))
	pp.Println(3)
	pp.Println("=========================================")
}

/**
 * 毎回到達可能であるかを調べる（10^4 オーダ）
 *   一番上のセルからスタートして、できるだけ下に進む
 *   検証にあたって、毎回検証用のmatrixをコピーする（10^4 オーダ）
 *   到達不可能であるとわかったセルは海にしていく
 *   到達できなくなったら結果を返却する
 *
 * セルがwaterになるごとに
 *   周辺4つのセルの到達可能性を下げる
 */

func latestDayToCross(row int, col int, cells [][]int) int {
	matrix := initMatrix(row, col)

	var lastDay int
	for day := range cells {
		r := cells[day][0] - 1
		c := cells[day][1] - 1
		matrix[r][c] = 0

		if !canCross(newMatrix(matrix)) {
			lastDay = day - 1
			break
		}
	}

	return lastDay + 1
}

func canCross(matrix [][]int) bool {
	var dfs func(row, col int) bool
	dfs = func(row, col int) bool {

		// reached
		if row == len(matrix)-1 {
			return true
		}

		matrix[row][col] = 0

		if row+1 != len(matrix) && matrix[row+1][col] == 1 && dfs(row+1, col) {
			return true
		}

		if col+1 != len(matrix[0]) && matrix[row][col+1] == 1 && dfs(row, col+1) {
			return true
		}

		if col != 0 && matrix[row][col-1] == 1 && dfs(row, col-1) {
			return true
		}

		return row != 0 && matrix[row-1][col] == 1 && dfs(row-1, col)
	}

	for col := range matrix[0] {
		if matrix[0][col] == 1 && dfs(0, col) {
			return true
		}
	}
	return false
}

func initMatrix(row, col int) [][]int {
	matrix := make([][]int, row)
	for r := range matrix {
		matrix[r] = make([]int, col)
		for c := range matrix[r] {
			matrix[r][c] = 1
		}
	}
	return matrix
}

func newMatrix(matrix [][]int) [][]int {
	newMatrix := make([][]int, len(matrix))
	for r := range newMatrix {
		newMatrix[r] = make([]int, len(matrix[0]))
		for c := range newMatrix[r] {
			newMatrix[r][c] = matrix[r][c]
		}
	}
	return newMatrix
}
