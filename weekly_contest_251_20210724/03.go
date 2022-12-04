// 5825. Maximum Compatibility Score Sum
// https://leetcode.com/contest/weekly-contest-251/problems/maximum-compatibility-score-sum/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(maxCompatibilitySum([][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 1}}, [][]int{{1, 0, 0}, {0, 0, 1}, {1, 1, 0}}))
	pp.Println(8)
	pp.Println("=========================================")
	pp.Println(maxCompatibilitySum([][]int{{0, 0}, {0, 0}, {0, 0}}, [][]int{{1, 1}, {1, 1}, {1, 1}}))
	pp.Println(0)
	pp.Println("=========================================")
	pp.Println(maxCompatibilitySum([][]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}, [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 1, 0},
	}))
	pp.Println(59)
	pp.Println("=========================================")
}

/**
 * 全パターンやって良さそう
 * パーミュテーションを作って
 *   8!
 * それぞれの組み合わせに対して
 *   * 8
 * スコアを計算して
 *   * 8
 * 最大値を保持しておく
 */
func maxCompatibilitySum(students [][]int, mentors [][]int) int {
	n, m := len(students[0]), len(students)

	var maxScore int

	var perm func(p []int, index int, used []bool)
	perm = func(p []int, index int, used []bool) {
		if index == m {
			var score int
			for i := 0; i < m; i++ {
				for j := 0; j < n; j++ {
					if students[i][j] == mentors[p[i]][j] {
						score += 1
					}
				}
			}
			if score > maxScore {
				maxScore = score
			}
			return
		}

		for i := 0; i < m; i++ {
			if !used[i] {
				p[index] = i
				used[i] = true
				perm(p, index+1, used)
				p[index] = i
				used[i] = false
			}
		}
	}

	perm(make([]int, m), 0, make([]bool, m))

	return maxScore
}
