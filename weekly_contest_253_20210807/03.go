// 5840. Minimum Number of Swaps to Make the String Balanced
// https://leetcode.com/contest/weekly-contest-253/problems/minimum-number-of-swaps-to-make-the-string-balanced/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(minSwaps("][]["))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(minSwaps("]]][[["))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(minSwaps("[]"))
	pp.Println(0)
	pp.Println("=========================================")
}

/**
 * Blanced string
 *   これまで開いた数よりも多く閉じてはいけない
 * 文字列を左から走査して
 *   閉じすぎているやつを見つけたら
 *   一番右にある開きすぎているやつと交換
 *   左右のポインタが出会ったら終了
 */

/**
 * refactored
 */
func minSwaps(s string) int {
	bs := []byte(s)
	l, r := 0, len(bs)-1
	var balance, count int
	for l <= r {
		if balance >= 0 {
			if bs[l] == '[' {
				balance++
			} else {
				balance--
			}
			l++
		} else if bs[r] != '[' {
			r--
		} else {
			bs[l-1], bs[r] = '[', ']'
			balance += 2
			count++
		}
	}
	return count
}

/**
 * Original Answer
 */
// func minSwaps(s string) int {
// 	arr := []byte(s)
// 	left, right := 0, len(arr)-1
// 	opened := 0
// 	count := 0
// 	for left < right {
// 	LOOP:
// 		for left < right {
// 			if arr[left] == '[' {
// 				opened++
// 				left++
// 			} else {
// 				if opened == 0 {
// 					break LOOP
// 				}
// 				opened--
// 				left++
// 			}
// 		}
//
// 		for arr[right] != '[' && left < right {
// 			right--
// 		}
//
// 		if left < right {
// 			arr[left], arr[right] = arr[right], arr[left]
// 			count++
// 		}
// 	}
// 	return count
// }

/**
 * ref: https://leetcode.com/problems/minimum-number-of-swaps-to-make-the-string-balanced/discuss/1390193/Python-Greedy-O(n)-explained
 */
// func minSwaps(s string) int {
// 	var maxUnbalance, curUnbalance int
// 	for _, b := range s {
// 		if b == ']' {
// 			curUnbalance++
// 		} else {
// 			curUnbalance--
// 		}
// 		maxUnbalance = max(maxUnbalance, curUnbalance)
// 	}
// 	return maxUnbalance>>1 + maxUnbalance&1
// }
//
// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }
