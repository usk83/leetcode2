// 5841. Find the Longest Valid Obstacle Course at Each Position
// https://leetcode.com/contest/weekly-contest-253/problems/find-the-longest-valid-obstacle-course-at-each-position/
package main

import (
	"fmt"
	"sort"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	fmt.Println(longestObstacleCourseAtEachPosition([]int{1, 2, 3, 2}))
	fmt.Println([]int{1, 2, 3, 3})
	pp.Println("=========================================")
	fmt.Println(longestObstacleCourseAtEachPosition([]int{2, 2, 1}))
	fmt.Println([]int{1, 2, 1})
	pp.Println("=========================================")
	fmt.Println(longestObstacleCourseAtEachPosition([]int{3, 1, 5, 6, 4, 2}))
	fmt.Println([]int{1, 1, 2, 3, 2, 2})
	pp.Println("=========================================")
}

func longestObstacleCourseAtEachPosition(obstacles []int) []int {
	n := len(obstacles)
	prevs := make([]int, 0, n)
	length := make([]int, n)
	for i := 0; i < n; i++ {
		index := sort.SearchInts(prevs, obstacles[i]+1)
		if index >= len(prevs) {
			prevs = append(prevs, obstacles[i])
		} else {
			prevs[index] = obstacles[i]
		}
		length[i] = index + 1
	}
	return length
}

/**
 * 自分より前で
 *   自分と同じか自分より小さくて
 *     そこまでの長さが一番長いものを選択する
 *   自分より大きい物しか存在しない場合もある
 *     その場合答えは1
 *
 * 2つの条件でソートされたものを取得したい
 *   障害物の高さ
 *   そこまでのコースの長さ
 *
 *     n == obstacles.length
 *     1 <= n <= 10^5
 *     1 <= obstacles[i] <= 10^7
 *
 * nlogn くらいまでの計算オーダ
 * 事前処理はOK
 *
 * 同じ長さのコースならより右のものを選択してしまってよい
 *
 * 位置ではソートされている
 * 高さでソートする
 * コースの長さでソートする
 *
 * Binary Search
 *
 * 位置とコースの長さ
 * コースの長さはindex
 */
// func longestObstacleCourseAtEachPosition(obstacles []int) []int {
// 	n := len(obstacles)
//
// 	// 位置と値の組み合わせを作って
// 	calcOrder := make([][2]int, n)
// 	for i, height := range obstacles {
// 		calcOrder[i] = [2]int{i, height}
// 	}
//
// 	// ソート
// 	// 値が小さいものが左
// 	// 同じ値のときは位置が左のものが左
// 	sort.Slice(calcOrder, func(i, j int) bool {
// 		return calcOrder[i][1] < calcOrder[j][1] || (calcOrder[i][1] == calcOrder[j][1] && calcOrder[i][0] < calcOrder[j][0])
// 	})
//
// 	// 位置とコースの長さ
// 	// コースの長さはindex
// 	optList := make([]int, 0, len(obstacles))
//
// 	length := make([]int, n)
// 	for i := o; i < n; i++ {
// 		// 自分より位置が左のもので一番右にあるものを探す
// 		i := sort.Search(len(optList), func(i int) bool {
//
// 		})
// 		calcOrder[0]
//
// 	}
// 	return length
// }
