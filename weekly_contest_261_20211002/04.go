// 5893. Smallest K-Length Subsequence With Occurrences of a Letter
// https://leetcode.com/contest/weekly-contest-261/problems/smallest-k-length-subsequence-with-occurrences-of-a-letter/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(smallestSubsequence("leet", 3, 'e', 1))
	pp.Println("eet")
	pp.Println("=========================================")
	pp.Println(smallestSubsequence("leetcode", 4, 'e', 2))
	pp.Println("ecde")
	pp.Println("=========================================")
	pp.Println(smallestSubsequence("bb", 2, 'b', 2))
	pp.Println("bb")
	pp.Println("=========================================")
}

/**
 * Constraints:
 * - 1 <= repetition <= k <= s.length <= 5 * 10^4
 * - s consists of lowercase English letters.
 * - letter is a lowercase English letter, and appears in s at least repetition times.
 */

/**
 * ref: https://leetcode.com/problems/smallest-k-length-subsequence-with-occurrences-of-a-letter/discuss/1500174/PythonJava-O(N)-greedy-solution-using-a-stack
 */
func smallestSubsequence(s string, requiredLength int, requiredLetter byte, requiredCount int) string {
	var remainRequiredLetterCount int
	for index := range s {
		letter := s[index]
		if letter == requiredLetter {
			remainRequiredLetterCount++
		}
	}

	var stack []byte

	for index := range s {
		letter := s[index]
		if stackLength := len(stack); stackLength != 0 {
			lastLetter := stack[stackLength-1]
			remainLetterCount := len(s) - index
			for stackLength != 0 && // スタックがあり
				lastLetter > letter && // スタックの末尾が現在の文字より大きく
				remainLetterCount+stackLength > requiredLength && // 残りの文字数とスタック内の文字数の合計が必要文字数より大きく
				(lastLetter != requiredLetter || remainRequiredLetterCount > requiredCount) { // 最後が必要な文字ではないか、そうだとしてもまだ必要文字数より必要文字の残り数が大きい
				stack = stack[:stackLength-1] // pop
				stackLength--
				if lastLetter == requiredLetter {
					requiredCount++
				}
			}
		}

		remainStackLength := len(stack)

		if remainStackLength < requiredLength {
			if letter == requiredLetter {
				stack = append(stack, letter)
				requiredCount--
			} else if requiredLength-remainStackLength > requiredCount {
				stack = append(stack, letter)
			}
		}

		if letter == requiredLetter {
			remainRequiredLetterCount--
		}
	}

	return string(stack)
}
