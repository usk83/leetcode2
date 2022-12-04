// 1897. Redistribute Characters to Make All Strings Equal
// https://leetcode.com/problems/redistribute-characters-to-make-all-strings-equal/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(makeEqual([]string{"abc", "aabc", "bc"}))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(makeEqual([]string{"ab", "a"}))
	pp.Println(false)
	pp.Println("=========================================")
}

/**
 * Runtime: 0 ms, faster than 100.00%
 * Memory Usage: 3.9 MB, less than 100.00%
 */
func makeEqual(words []string) bool {
	var charCount [26]int
	for _, word := range words {
		for _, char := range word {
			charCount[char-'a']++
		}
	}
	for _, count := range charCount {
		if count%len(words) != 0 {
			return false
		}
	}
	return true
}
