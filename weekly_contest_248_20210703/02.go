// 5801. Eliminate Maximum Number of Monsters
// https://leetcode.com/contest/weekly-contest-248/problems/eliminate-maximum-number-of-monsters/
package main

import (
	"sort"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(eliminateMaximum([]int{1, 3, 4}, []int{1, 1, 1}))
	pp.Println(3)
	pp.Println("=========================================")
	pp.Println(eliminateMaximum([]int{1, 1, 2, 3}, []int{1, 1, 1, 1}))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(eliminateMaximum([]int{3, 2, 4}, []int{5, 3, 2}))
	pp.Println(1)
	pp.Println("=========================================")
}

/**
 * それぞれがどれだけで到着するか計算して
 * ソートして
 * 頭から比較して、今の時間が超えたら（同じも）アウト
 */
func eliminateMaximum(dist []int, speed []int) int {
	n := len(dist)
	time := make([]int, n)
	for i := range time {
		time[i] = dist[i] / speed[i]
		if dist[i]%speed[i] != 0 {
			time[i]++
		}
	}

	sort.Ints(time)

	var count int
	for i := 0; i < n; i++ {
		if i >= time[i] {
			break
		}
		count++
	}
	return count
}
