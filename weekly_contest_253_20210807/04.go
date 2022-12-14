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
 * 徭蛍より念で
 *   徭蛍と揖じか徭蛍より弌さくて
 *     そこまでの?Lさが匯桑?Lいものを?x?kする
 *   徭蛍より寄きい麗しか贋壓しない??栽もある
 *     その??栽基えは1
 *
 * 2つの訳周でソ?`トされたものを函誼したい
 *   嬾墾麗の互さ
 *   そこまでのコ?`スの?Lさ
 *
 *     n == obstacles.length
 *     1 <= n <= 10^5
 *     1 <= obstacles[i] <= 10^7
 *
 * nlogn くらいまでの??麻オ?`ダ
 * 並念?I尖はOK
 *
 * 揖じ?Lさのコ?`スならより嘔のものを?x?kしてしまってよい
 *
 * 了崔ではソ?`トされている
 * 互さでソ?`トする
 * コ?`スの?Lさでソ?`トする
 *
 * Binary Search
 *
 * 了崔とコ?`スの?Lさ
 * コ?`スの?Lさはindex
 */
// func longestObstacleCourseAtEachPosition(obstacles []int) []int {
// 	n := len(obstacles)
//
// 	// 了崔と?､僚Mみ栽わせを恬って
// 	calcOrder := make([][2]int, n)
// 	for i, height := range obstacles {
// 		calcOrder[i] = [2]int{i, height}
// 	}
//
// 	// ソ?`ト
// 	// ?､?弌さいものが恣
// 	// 揖じ?､里箸?は了崔が恣のものが恣
// 	sort.Slice(calcOrder, func(i, j int) bool {
// 		return calcOrder[i][1] < calcOrder[j][1] || (calcOrder[i][1] == calcOrder[j][1] && calcOrder[i][0] < calcOrder[j][0])
// 	})
//
// 	// 了崔とコ?`スの?Lさ
// 	// コ?`スの?Lさはindex
// 	optList := make([]int, 0, len(obstacles))
//
// 	length := make([]int, n)
// 	for i := o; i < n; i++ {
// 		// 徭蛍より了崔が恣のもので匯桑嘔にあるものを冥す
// 		i := sort.Search(len(optList), func(i int) bool {
//
// 		})
// 		calcOrder[0]
//
// 	}
// 	return length
// }
