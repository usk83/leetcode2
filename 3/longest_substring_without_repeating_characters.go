// 3. Longest Substring Without Repeating Characters
// https://leetcode.com/problems/longest-substring-without-repeating-characters/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(lengthOfLongestSubstring("abcabcbb"))
	pp.Println(3)
	pp.Println("=========================================")
	pp.Println(lengthOfLongestSubstring("bbbbb"))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(lengthOfLongestSubstring("pwwkew"))
	pp.Println(3)
	pp.Println("=========================================")
	pp.Println(lengthOfLongestSubstring(""))
	pp.Println(0)
	pp.Println("=========================================")
}

/**
 * Constraints:
 * - 0 <= s.length <= 5 * 10^4
 * - s consists of English letters, digits, symbols and spaces.
 *
 * N: Length of the given string
 */

/**
 * 2. Keep the index of the last appearance of each character using HashMap,
 *    and keep track of the left index of the current substring to calculate the length it
 * - Time Complexity: O(N)
 * - Space Complexity: O(N)
 * > Runtime: 4 ms (faster than 91.95%)
 * > Memory Usage: 3.2 MB (less than 48.95%)
 */
func lengthOfLongestSubstring(s string) int {
	// 使われている時
	// 	それが最後に出たところが今のheadよりも後ろならその次をheadにする
	// 最後の出現箇所を更新する
	// 	headから現在位置までの長さで最長を更新する
	charIndexMap, left, longest := map[rune]int{}, 0, 0
	for right, r := range s {
		if last, ok := charIndexMap[r]; ok && last >= left {
			left = last + 1
		}
		charIndexMap[r] = right
		longest = max(longest, right-left+1)
	}
	return longest
}

/**
 * 1. Keep a hashmap holding characters used in the current substring
 * - Time Complexity: O(N)
 * - Space Complexity: O(N)
 * > Runtime: 0 ms (faster than 100.00%)
 * > Memory Usage: 3.1 MB (less than 59.64%)
 */
// func lengthOfLongestSubstring(s string) int {
// 	// 使ってない限りは入れていく
// 	// 	入れたときに最長を更新する
// 	// 	最後のindexを消費するまで
// 	// 使われている時は、それを吐き出すまで左を落としていく
// 	rs, charSet, longest := []rune(s), map[rune]struct{}{}, 0
// 	for i, r := range rs {
// 		if _, ok := charSet[r]; ok {
// 			for rs[i-len(charSet)] != r {
// 				delete(charSet, rs[i-len(charSet)])
// 			}
// 		} else {
// 			charSet[r] = struct{}{}
// 		}
// 		longest = max(longest, len(charSet))
// 	}
// 	return longest
// }

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
