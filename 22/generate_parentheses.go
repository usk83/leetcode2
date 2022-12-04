// 22. Generate Parentheses
// https://leetcode.com/problems/generate-parentheses/
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	fmt.Println(generateParenthesis(3))
	fmt.Println([]string{"((()))", "(()())", "(())()", "()(())", "()()()"})
	pp.Println("=========================================")
	fmt.Println(generateParenthesis(1))
	fmt.Println([]string{"()"})
	pp.Println("=========================================")
}

/**
 * 1. DFS
 * - Time Complexity: O(4^N/√N) according to https://leetcode.com/problems/generate-parentheses/solution/
 * - Space Complexity: O(4^N/√N) according to https://leetcode.com/problems/generate-parentheses/solution/
 * > Runtime: 0 ms (faster than 100.00%)
 * > Memory Usage: 2.8 MB (less than 94.08%)
 */
func generateParenthesis(n int) []string {
	var combinations []string
	var generate func(int, int, []rune)
	generate = func(left, right int, parenthesis []rune) {
		if left == 0 && right == 0 {
			combinations = append(combinations, string(parenthesis))
			return
		}
		if left != 0 {
			generate(left-1, right, append(parenthesis, '('))
		}
		if right > left {
			generate(left, right-1, append(parenthesis, ')'))
		}
	}
	generate(n, n, make([]rune, 0, n<<1))
	return combinations
}
