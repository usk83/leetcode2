// 20. Valid Parentheses
// https://leetcode.com/problems/valid-parentheses/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(isValid("()"))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(isValid("()[]{}"))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(isValid("(]"))
	pp.Println(false)
	pp.Println("=========================================")
	pp.Println(isValid("([)]"))
	pp.Println(false)
	pp.Println("=========================================")
	pp.Println(isValid("{[]}"))
	pp.Println(true)
	pp.Println("=========================================")
}

/**
 * [description]
 * - Time Complexity: O(X)
 * - Space Complexity: O(X)
 */

var pairs = map[rune]rune{
	')': '(',
	'}': '{',
	']': '[',
}

// func isValid(s string) bool {
// 	stack := []rune{}
// 	for _, r := range s {
// 		if pair, ok := pairs[r]; ok {
// 			if len(stack) == 0 || stack[len(stack)-1] != pair {
// 				return false
// 			}
// 			stack = stack[:len(stack)-1] // pop
// 		} else {
// 			stack = append(stack, r) // push
// 		}
// 	}
// 	return len(stack) == 0
// }

func isValid(s string) bool {
	stack := []rune{}
	var count int
	for _, b := range s {
		switch b {
		case '(', '{', '[':
			stack = append(stack, b) // push
			goto label2
		label:
			pp.Println("debugging...1")
			count++
			if count > 3 {
				return false
			}
		label2:
			pp.Println("debugging...2")
			goto label
			break
		default:
			var pair rune
			switch b {
			case ')':
				pair = '('
			case '}':
				pair = '{'
			case ']':
				pair = '['
			}
			if len(stack) == 0 {
				return false
			}
			lastRune := stack[len(stack)-1]
			if lastRune != pair {
				return false
			}
			stack = stack[:len(stack)-1] // pop
		}
	}
	return len(stack) == 0
}

// func isValid(s string) bool {
// 	stack := []rune{}
// 	for _, b := range s {
// 		if b == '(' || b == '{' || b == '[' {
// 			stack = append(stack, b) // push
// 		} else {
// 			var pair rune
// 			if b == ')' {
// 				pair = '('
// 			} else if b == '}' {
// 				pair = '{'
// 			} else if b == ']' {
// 				pair = '['
// 			}
// 			if len(stack) == 0 {
// 				return false
// 			}
// 			lastRune := stack[len(stack)-1]
// 			if lastRune != pair {
// 				return false
// 			}
// 			stack = stack[:len(stack)-1] // pop
// 		}
// 	}
// 	return len(stack) == 0
// }
