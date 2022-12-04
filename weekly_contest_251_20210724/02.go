// 5824. Largest Number After Mutating Substring
// https://leetcode.com/contest/weekly-contest-251/problems/largest-number-after-mutating-substring/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(maximumNumber("132", []int{9, 8, 5, 0, 3, 6, 4, 2, 6, 8}))
	pp.Println("832")
	pp.Println("=========================================")
	pp.Println(maximumNumber("021", []int{9, 4, 3, 5, 7, 2, 1, 9, 0, 6}))
	pp.Println("934")
	pp.Println("=========================================")
	pp.Println(maximumNumber("5", []int{1, 4, 7, 5, 3, 2, 5, 6, 9, 4}))
	pp.Println("5")
	pp.Println("=========================================")
	pp.Println(maximumNumber("214010", []int{6, 7, 9, 7, 4, 0, 3, 4, 4, 7}))
	pp.Println("974676")
	pp.Println("=========================================")
	pp.Println(maximumNumber("334111", []int{0, 9, 2, 3, 3, 2, 5, 5, 5, 5}))
	pp.Println("334999")
	pp.Println("=========================================")
}

/**
 * 左から順に見て
 * 交換したほうがいいものがあったとき
 * そこから右に可能な限り交換する
 */
func maximumNumber(num string, change []int) string {
	state := 0 // 0: not, 1: converting, 2: finished
	var maxBytes []rune
	for _, b := range num {
		if state != 2 {
			cur := int(b - 48)
			if cur <= change[cur] {
				maxBytes = append(maxBytes, rune(change[cur]+48))
				if cur != change[cur] {
					state = 1
				}
			} else {
				maxBytes = append(maxBytes, b)
				if state == 1 {
					state = 2
				}
			}
		} else {
			maxBytes = append(maxBytes, b)
		}
	}
	return string(maxBytes)
}
