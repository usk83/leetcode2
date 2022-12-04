// 2000. Reverse Prefix of Word
// https://leetcode.com/problems/reverse-prefix-of-word/
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
 * Constraints:
 * - 1 <= word.length <= 250
 * - word consists of lowercase English letters.
 * - ch is a lowercase English letter.
 *
 * N: Length of the given `word`
 */

/**
 * 1. Find the first appearance of `ch` and reverse the former substring
 * - Time Complexity: O(N)
 * - Space Complexity: O(N)
 * > Runtime: 0 ms (faster than 100.00%)
 * > Memory Usage: 2.1 MB (less than 100.00%)
 */
func reversePrefix(word string, ch byte) string {
	for i := range word {
		if word[i] != ch {
			continue
		}
		bs := []byte(word)
		for j := i; j > i>>1; j-- {
			bs[j], bs[i-j] = bs[i-j], bs[j]
		}
		return string(bs)
	}
	return word
}
