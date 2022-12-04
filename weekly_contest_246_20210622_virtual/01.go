// 1903. Largest Odd Number in String
// https://leetcode.com/contest/weekly-contest-246/problems/largest-odd-number-in-string/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(largestOddNumber("52"))
	pp.Println("5")
	pp.Println("=========================================")
	pp.Println(largestOddNumber("4206"))
	pp.Println("")
	pp.Println("=========================================")
	pp.Println(largestOddNumber("35427"))
	pp.Println("35427")
	pp.Println("=========================================")
}

func largestOddNumber(num string) string {
	chars := []byte(num)
	lastOdd := len(chars) - 1
	for ; lastOdd >= 0; lastOdd-- {
		if (chars[lastOdd]-48)&1 == 1 {
			break
		}
	}
	return string(chars[0 : lastOdd+1])
}
