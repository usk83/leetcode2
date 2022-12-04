// 5864. The Number of Weak Characters in the Game
// https://leetcode.com/contest/weekly-contest-257/problems/the-number-of-weak-characters-in-the-game/
package main

import (
	"sort"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(numberOfWeakCharacters([][]int{{5, 5}, {6, 3}, {3, 6}}))
	pp.Println(0)
	pp.Println("=========================================")
	pp.Println(numberOfWeakCharacters([][]int{{2, 2}, {3, 3}}))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(numberOfWeakCharacters([][]int{{1, 5}, {10, 4}, {4, 3}}))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(numberOfWeakCharacters([][]int{{1, 1}, {2, 1}, {2, 2}, {1, 2}}))
	pp.Println(1)
	pp.Println("=========================================")
}

/**
 * [description]
 * - Time Complexity: O(X)
 * - Space Complexity: O(X)
 * > Runtime: X ms (faster than X%)
 * > Memory Usage: X MB (less than X%)
 */

// 2 <= properties.length <= 10^5
// properties[i].length == 2
// 1 <= attacki, defensei <= 10^5
func numberOfWeakCharacters(properties [][]int) int {
	sort.Slice(properties, func(i, j int) bool {
		return properties[i][0] > properties[j][0] || (properties[i][0] == properties[j][0] && properties[i][1] > properties[j][1])
	})

	var lastAtk, tmpMaxDef, maxDef, count int
	for _, prop := range properties {
		if prop[0] != lastAtk {
			maxDef = tmpMaxDef
		}
		if prop[1] < maxDef {
			count++
		} else {
			tmpMaxDef = max(tmpMaxDef, prop[1])
		}
		lastAtk = prop[0]
	}
	return count
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
