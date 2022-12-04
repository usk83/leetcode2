// 5890. Minimum Moves to Convert String
// https://leetcode.com/contest/weekly-contest-261/problems/minimum-moves-to-convert-string/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(minimumMoves("XXX"))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(minimumMoves("XXOX"))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(minimumMoves("OOOO"))
	pp.Println(0)
	pp.Println("=========================================")
}

/**
 * Constraints:
 * - 3 <= s.length <= 1000
 * - s[i] is either 'X' or 'O'.
 *
 * N: Length of the given string
 */

/**
 * 2. Efficient solution only requires constant space
 * - Time Complexity: O(N)
 * - Space Complexity: O(1)
 * > Runtime: 0 ms (faster than 100.00%)
 * > Memory Usage: 2.2 MB (less than 100.00%)
 */
func minimumMoves(s string) int {
	var count int
	for i := 0; i < len(s); i++ {
		if s[i] == 'X' {
			count++
			i += 2
		}
	}
	return count
}

/**
 * 1. Original solution
 * - Time Complexity: O(N)
 * - Space Complexity: O(N)
 */
// func minimumMoves(s string) int {
// 	min := func(x, y int) int {
// 		if x < y {
// 			return x
// 		}
// 		return y
// 	}
//
// 	arr := make([]int, len(s))
// 	for i := range s {
// 		if s[i] == 'X' {
// 			arr[i] = 1
// 		}
// 	}
// 	var count int
// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] == 1 {
// 			count++
// 			arr[i] = 0
// 			arr[min(len(s)-1, i+1)] = 0
// 			arr[min(len(s)-1, i+2)] = 0
// 		}
// 	}
// 	return count
// }
