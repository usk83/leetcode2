// 5802. Count Good Numbers
// https://leetcode.com/contest/weekly-contest-248/problems/count-good-numbers/
package main

import "github.com/k0kubun/pp"

func main() {
	pp.Println("=========================================")
	pp.Println(countGoodNumbers(1))
	pp.Println(5)
	pp.Println("=========================================")
	pp.Println(countGoodNumbers(4))
	pp.Println(400)
	pp.Println("=========================================")
	pp.Println(countGoodNumbers(50))
	pp.Println(564908303)
	pp.Println("=========================================")
}

/**
 * 偶数番目であることが許されるのは
 * 0, 2, 4, 6, 8
 *
 * 奇数番目であることが許されるのは
 * 2, 3, 5, 7
 *
 * 2の何乗で示すことができるかを考えて
 * 計算を圧縮する
 */

const MOD = 1000000007

// 1 <= n <= 10^15
func countGoodNumbers(n int64) int {
	evens := int(n>>1 + n&1)
	odds := int(n >> 1)

	count := 1
	count = evenCalc(count, evens)
	count = oddCalc(count, odds)

	return count
}

func evenCalc(num, c int) int {
	for c >= 2 {
		p := 1
		nn := 2
		for nn<<1 <= c {
			nn <<= 1
			p++
		}
		c -= nn

		mul := 5
		for i := 0; i < p; i++ {
			mul = mul * mul % MOD
		}

		num = num * mul % MOD
	}

	if c == 1 {
		num = num * 5 % MOD
	}

	return num
}

func oddCalc(num, c int) int {
	for c >= 2 {
		p := 1
		nn := 2
		for nn<<1 <= c {
			nn <<= 1
			p++
		}
		c -= nn

		mul := 4
		for i := 0; i < p; i++ {
			mul = mul * mul % MOD
		}

		num = num * mul % MOD
	}

	if c > 1 {
		panic("")
	}

	if c == 1 {
		num = num * 4 % MOD
	}

	return num
}
