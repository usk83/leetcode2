// 5838. Check If String Is a Prefix of Array
// https://leetcode.com/contest/weekly-contest-253/problems/check-if-string-is-a-prefix-of-array/
package main

import (
	"strings"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(isPrefixString("iloveleetcode", []string{"i", "love", "leetcode", "apples"}))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(isPrefixString("iloveleetcode", []string{"apples", "i", "love", "leetcode"}))
	pp.Println(false)
	pp.Println("=========================================")
}

func isPrefixString(s string, words []string) bool {
	index := 0
	for s != "" {
		if index == len(words) {
			return false
		}
		if strings.HasPrefix(s, words[index]) {
			s = string(s[len(words[index]):])
			index++
		} else {
			return false
		}
	}
	return true
}
