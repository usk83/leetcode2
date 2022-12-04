// 5952. Rings and Rods
// https://leetcode.com/contest/weekly-contest-271/problems/rings-and-rods/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(countPoints("B0B6G0R6R0R6G9"))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(countPoints("B0R0G0R9R0B0G0"))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(countPoints("G4"))
	pp.Println(0)
	pp.Println("=========================================")
	pp.Println(countPoints("R0G0B0R1G2B3"))
	pp.Println(1)
	pp.Println("=========================================")
}

func countPoints(rings string) int {
	rods := [10][3]int{}
	for i := 0; i < len(rings)>>1; i++ {
		switch rings[i<<1] {
		case 'R':
			rods[rings[i<<1+1]-'0'][0] = 1
		case 'G':
			rods[rings[i<<1+1]-'0'][1] = 1
		case 'B':
			rods[rings[i<<1+1]-'0'][2] = 1
		default:
		}
	}

	var count int
	for _, rod := range rods {
		if rod[0] == 1 && rod[1] == 1 && rod[2] == 1 {
			count++
		}
	}
	return count
}
