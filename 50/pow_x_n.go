// 50. Pow(x, n)
// https://leetcode.com/problems/powx-n/
package main

import (
	"math/bits"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(myPow(2.00000, 10))
	pp.Println(1024.00000)
	pp.Println("=========================================")
	pp.Println(myPow(2.10000, 3))
	pp.Println(9.26100)
	pp.Println("=========================================")
	pp.Println(myPow(34.00515, -3))
	pp.Println(3e-05)
	pp.Println("=========================================")
}

/**
 * Constraints:
 * - -100.0 < x < 100.0
 * - -2^31 <= n <= 2^31 - 1
 * - -10^4 <= x^n <= 10^4
 */

/**
 * 2-4. Smartest iteration
 * ref: https://leetcode.com/problems/powx-n/discuss/19560/Shortest-Python-Guaranteed
 * - Time Complexity: O(logN)
 * - Space Complexity: O(logN)
 */
// func myPow(x float64, n int) float64 {
// 	if n < 0 {
// 		x = 1 / x
// 		n *= -1
// 	}
// 	var pow float64 = 1.0
// 	for n != 0 {
// 		if n&1 == 1 {
// 			pow *= x
// 		}
// 		x *= x
// 		n >>= 1
// 	}
// 	return pow
// }

/**
 * 2-3. Refactored smarter iteration
 * - Time Complexity: O(logN)
 * - Space Complexity: O(logN)
 * > Runtime: 0 ms (faster than 100.00%)
 * > Memory Usage: 2.1 MB (less than 100.00%)
 */
func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n *= -1
	}
	var pow float64 = 1.0
	for b := bits.Len64(uint64(n)) - 1; b >= 0; b-- {
		pow *= pow
		if n>>b&1 == 1 {
			pow *= x
		}
	}
	return pow
}

/**
 * 2-2. Smarter iteration using bitwise operations
 * - Time Complexity: O(logN)
 * - Space Complexity: O(logN)
 */
// func myPow(x float64, n int) float64 {
// 	isMinus := n < 0
// 	if isMinus {
// 		n *= -1
// 	}
// 	var pow float64 = 1.0
// 	for b := bits.Len64(uint64(n)) - 1; b >= 0; b-- {
// 		pow = pow * pow
// 		if n>>b&1 == 1 {
// 			pow *= x
// 		}
// 	}
// 	if isMinus {
// 		return 1 / pow
// 	}
// 	return pow
// }

/**
 * 2-1. Iteration
 * - Time Complexity: O(logN)
 * - Space Complexity: O(logN)
 */
// func myPow(x float64, n int) float64 {
// 	isMinus := n < 0
// 	if isMinus {
// 		n *= -1
// 	}
// 	breakdown := []bool{}
// 	for ; n != 0; n >>= 1 {
// 		breakdown = append(breakdown, n&1 == 1)
// 	}
// 	var pow float64 = 1.0
// 	for i := len(breakdown) - 1; i >= 0; i-- {
// 		pow = pow * pow
// 		if breakdown[i] {
// 			pow *= x
// 		}
// 	}
// 	if isMinus {
// 		return 1 / pow
// 	}
// 	return pow
// }

/**
 * 1-2. Refactored recursion
 * - Time Complexity: O(logN)
 * - Space Complexity: O(logN)
 * > Runtime: 0 ms (faster than 100.00%)
 * > Memory Usage: 2.1 MB (less than 100.00%)
 */
// func myPow(x float64, n int) float64 {
// 	switch n {
// 	case -1:
// 		return 1 / x
// 	case 0:
// 		return 1
// 	case 1:
// 		return x
// 	}
// 	return myPow(x*x, n/2) * myPow(x, n%2)
// }

/**
 * 1-1. Recursion
 * - Time Complexity: O(logN)
 * - Space Complexity: O(logN)
 * > Runtime: 0 ms (faster than 100.00%)
 * > Memory Usage: 2.2 MB (less than 41.96%)
 */
// func myPow(x float64, n int) float64 {
// 	switch n {
// 	case -1:
// 		return 1 / x
// 	case 0:
// 		return 1
// 	case 1:
// 		return x
// 	}
// 	halfPow, mod := myPow(x, n/2), n%2
// 	return halfPow * halfPow * myPow(x, mod)
// }

/**
 * 0. Naive recursion
 * - Time Complexity: O(N)
 * - Space Complexity: O(N)
 */
// func myPow(x float64, n int) float64 {
// 	if n > 0 {
// 		return x * myPow(x, n-1)
// 	} else if n < 0 {
// 		return 1 / x * myPow(x, n+1)
// 	}
// 	return 1
// }
