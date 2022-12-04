// 5892. Stone Game IX
// https://leetcode.com/contest/weekly-contest-261/problems/stone-game-ix/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(stoneGameIX([]int{2, 1}))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(stoneGameIX([]int{2}))
	pp.Println(false)
	pp.Println("=========================================")
	pp.Println(stoneGameIX([]int{5, 1, 2, 4, 3}))
	pp.Println(false)
	pp.Println("=========================================")
	pp.Println(stoneGameIX([]int{3}))
	pp.Println(false)
	pp.Println("=========================================")
	pp.Println(stoneGameIX([]int{2, 2, 3}))
	pp.Println(false)
	pp.Println("=========================================")
	pp.Println(stoneGameIX([]int{2, 3, 1}))
	pp.Println(false)
	pp.Println("=========================================")
	pp.Println(stoneGameIX([]int{20, 3, 20, 17, 2, 12, 15, 17, 4}))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(stoneGameIX([]int{15, 20, 10, 13, 14, 15, 5, 2, 3}))
	pp.Println(false)
	pp.Println("=========================================")
}

/**
 * Constraints:
 * - 1 <= stones.length <= 10^5
 * - 1 <= stones[i] <= 10^4
 *
 * N: Number of stones
 */

/**
 * 1. Original solution
 * - Time Complexity: O(N)
 * - Space Complexity: O(1)
 * > Runtime: 188 ms (faster than 100.00%)
 * > Memory Usage: 9.1 MB (less than 100.00%)
 */
func stoneGameIX(rawStones []int) bool {
	var stones [3]int
	for _, rawStone := range rawStones {
		stones[rawStone%3]++
	}
	if stones[0]%2 == 0 {
		return stones[1] != 0 && stones[2] != 0
	}
	sub := stones[1] - stones[2]
	return sub <= -3 || 3 <= sub
}
