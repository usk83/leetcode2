// 5843. Number of Strings That Appear as Substrings in Word
// https://leetcode.com/contest/weekly-contest-254/problems/number-of-strings-that-appear-as-substrings-in-word/
package main

import (
	"strings"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
}

func numOfStrings(patterns []string, word string) int {
	var count int
	for _, p := range patterns {
		if strings.Contains(word, p) {
			count++
		}
	}
	return count
}
