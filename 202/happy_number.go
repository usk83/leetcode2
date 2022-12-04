// 202. Happy Number
// https://leetcode.com/problems/happy-number/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(isHappy(19))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(isHappy(2))
	pp.Println(false)
	pp.Println("=========================================")
}

/**
 * v2. Two Pointers
 * - Time Complexity: O(logN)
 * - Space Complexity: O(1)
 */
func isHappy(n int) bool {
	proc := func(n int) int {
		var next, mod int
		for n != 0 {
			n, mod = n/10, n%10
			next += mod * mod
		}
		return next
	}
	slow, fast := n, proc(n)
	for slow != fast {
		slow, fast = proc(slow), proc(proc(fast))
	}
	return slow == 1
}

/**
 * v1. Set
 * - Time Complexity: O(logN)
 * - Space Complexity: O(logN)
 */
// func isHappy(n int) bool {
// 	set := map[int]bool{1: true}
// 	for !set[n] {
// 		set[n] = true
// 		var oldN int
// 		for n, oldN = 0, n; oldN != 0; oldN = oldN / 10 {
// 			mod := oldN % 10
// 			n += mod * mod
// 		}
// 	}
// 	return n == 1
// }
