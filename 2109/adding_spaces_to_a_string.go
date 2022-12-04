// 2109. Adding Spaces to a String
// https://leetcode.com/problems/adding-spaces-to-a-string/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(addSpaces("LeetcodeHelpsMeLearn", []int{8, 13, 15}))
	pp.Println("Leetcode Helps Me Learn")
	pp.Println("=========================================")
	pp.Println(addSpaces("icodeinpython", []int{1, 5, 7, 9}))
	pp.Println("i code in py thon")
	pp.Println("=========================================")
	pp.Println(addSpaces("spacing", []int{0, 1, 2, 3, 4, 5, 6}))
	pp.Println(" s p a c i n g")
	pp.Println("=========================================")
}

/**
 * Constraints:
 * - 1 <= s.length <= 3 * 10^5
 * - s consists only of lowercase and uppercase English letters.
 * - 1 <= spaces.length <= 3 * 10^5
 * - 0 <= spaces[i] <= s.length - 1
 * - All the values of spaces are strictly increasing.
 *
 * N: Length of the given string
 * M: Number of spaces to add
 */

/**
 * 1. Copy parts of the original string and add a space one by one
 * - Time Complexity: O(N+M)
 * - Space Complexity: O(N+M)
 * > Runtime: 100 ms (faster than 100.00%)
 * > Memory Usage: 14 MB (less than 89.19%)
 */
func addSpaces(s string, spaces []int) string {
	numOfSpaces := len(spaces)
	ss := make([]byte, len(s)+numOfSpaces)
	bs := []byte(s)
	var head int
	for i, idx := range spaces {
		copy(ss[head+i:], bs[head:idx])
		ss[idx+i] = ' '
		head = idx
	}
	lastSpaceIndex := spaces[numOfSpaces-1]
	copy(ss[lastSpaceIndex+numOfSpaces:], bs[lastSpaceIndex:])
	return string(ss)
}
