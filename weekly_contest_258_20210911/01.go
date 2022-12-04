// 5867. Reverse Prefix of Word
// https://leetcode.com/contest/weekly-contest-258/problems/reverse-prefix-of-word/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(reversePrefix("abcdefd", 'd'))
	pp.Println("dcbaefd")
	pp.Println("=========================================")
	pp.Println(reversePrefix("xyxzxe", 'z'))
	pp.Println("zxyxxe")
	pp.Println("=========================================")
	pp.Println(reversePrefix("abcd", 'z'))
	pp.Println("abcd")
	pp.Println("=========================================")
}

/**
 * 1. Original solution
 * - Time Complexity: O(N)
 * - Space Complexity: O(N)
 */
func reversePrefix(word string, ch byte) string {
	bs := []byte(word)
	index := -1
	for i := range bs {
		if bs[i] == ch {
			index = i
			break
		}
	}
	if index == -1 {
		return word
	}

	ret := []byte{}
	for i := index; i >= 0; i-- {
		ret = append(ret, bs[i])
	}
	for i := index + 1; i < len(bs); i++ {
		ret = append(ret, bs[i])
	}
	return string(ret)
}
