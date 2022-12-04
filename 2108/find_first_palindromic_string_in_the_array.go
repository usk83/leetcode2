// 2108. Find First Palindromic String in the Array
// https://leetcode.com/problems/find-first-palindromic-string-in-the-array/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(firstPalindrome([]string{"abc", "car", "ada", "racecar", "cool"}))
	pp.Println("ada")
	pp.Println("=========================================")
	pp.Println(firstPalindrome([]string{"notapalindrome", "racecar"}))
	pp.Println("racecar")
	pp.Println("=========================================")
	pp.Println(firstPalindrome([]string{"def", "ghi"}))
	pp.Println("")
	pp.Println("=========================================")
}

/**
 * Constraints:
 * - 1 <= words.length <= 100
 * - 1 <= words[i].length <= 100
 * - words[i] consists only of lowercase English letters.
 *
 * N: Number of the given `words`
 * M: Total number of characters of given words
 */

/**
 * 2. Define a function to check if a string is palindromic or not
 * - Time Complexity: O(M)
 * - Space Complexity: O(1)
 * > Runtime: 16 ms (faster than 81.58%)
 * > Memory Usage: 6.7 MB (less than 86.84%)
 */
func firstPalindrome(words []string) string {
	for _, word := range words {
		if isPalindrome(word) {
			return word
		}
	}
	return ""
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

/**
 * 1. Inline check if a word is palindromic or not (faster)
 * - Time Complexity: O(M)
 * - Space Complexity: O(1)
 * > Runtime: 4 ms (faster than 100.00%)
 * > Memory Usage: 6.7 MB (less than 86.84%)
 */
// func firstPalindrome(words []string) string {
// LOOP:
// 	for _, word := range words {
// 		for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
// 			if word[i] != word[j] {
// 				continue LOOP
// 			}
// 		}
// 		return word
// 	}
// 	return ""
// }
