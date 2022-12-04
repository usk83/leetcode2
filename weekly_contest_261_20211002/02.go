// 5891. Find Missing Observations
// https://leetcode.com/contest/weekly-contest-261/problems/find-missing-observations/
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	fmt.Println(missingRolls([]int{3, 2, 4, 3}, 4, 2))
	fmt.Println([]int{6, 6})
	pp.Println("=========================================")
	fmt.Println(missingRolls([]int{1, 5, 6}, 3, 4))
	fmt.Println([]int{2, 3, 2, 2})
	pp.Println("=========================================")
	fmt.Println(missingRolls([]int{1, 2, 3, 4}, 6, 4))
	fmt.Println([]int{})
	pp.Println("=========================================")
	fmt.Println(missingRolls([]int{1}, 3, 1))
	fmt.Println([]int{5})
	pp.Println("=========================================")
}

/**
 * Constraints:
 * - m == rolls.length
 * - 1 <= n, m <= 10^5
 * - 1 <= rolls[i], mean <= 6
 *
 * N: Number of missssing dice rolls
 * M: Number of observable dice rolls
 */

/**
 * 2. Refactored solution
 * - Time Complexity: O(N+M)
 * - Space Complexity: O(N)
 * > Runtime: 172 ms (faster than 100.00%)
 * > Memory Usage: 9.4 MB (less than 100.00%)
 */
func missingRolls(rolls []int, mean int, n int) []int {
	m := len(rolls)
	sum := mean * (n + m)
	for _, roll := range rolls {
		sum -= roll
	}
	if sum < n || n*6 < sum {
		return []int{}
	}
	missings, extra := make([]int, n), sum-n
	for i := 0; i < n; i++ {
		if extra == 0 {
			missings[i] = 1
		} else if extra < 6 {
			missings[i] = 1 + extra
			extra = 0
		} else {
			missings[i] = 6
			extra -= 5
		}
	}
	return missings
}

/**
 * 1. Original solution
 * - Time Complexity: O(N+M)
 * - Space Complexity: O(N)
 */
// func missingRolls(rolls []int, mean int, n int) []int {
// 	// ない場合がある
// 	// 	空配列を返す
// 	m := len(rolls)
// 	sum := mean * (n + m)
// 	for _, roll := range rolls {
// 		sum -= roll
// 	}
// 	if sum < n || sum > n*6 {
// 		return []int{}
// 	}
// 	ans := make([]int, n)
// 	for i := range ans {
// 		ans[i] = 1
// 	}
// 	sum -= n
// 	index := 0
// 	for sum != 0 {
// 		if sum >= 5 {
// 			ans[index] = 6
// 			index++
// 			sum -= 5
// 		} else {
// 			ans[index] += sum
// 			sum = 0
// 		}
// 	}
// 	return ans
// }
