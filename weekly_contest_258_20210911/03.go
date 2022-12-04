// 5869. Maximum Product of the Length of Two Palindromic Subsequences
// https://leetcode.com/contest/weekly-contest-258/problems/maximum-product-of-the-length-of-two-palindromic-subsequences/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(maxProduct("leetcodecom"))
	pp.Println(9)
	pp.Println("=========================================")
	pp.Println(maxProduct("bb"))
	pp.Println(1)
	pp.Println("=========================================")
	pp.Println(maxProduct("accbcaxxcxx"))
	pp.Println(25)
	pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
}

/**
 * [description]
 * - Time Complexity: O(X)
 * - Space Complexity: O(X)
 * > Runtime: X ms (faster than X%)
 * > Memory Usage: X MB (less than X%)
 */

// 2 <= s.length <= 12
// s consists of lowercase English letters only.
func maxProduct(s string) int {
	ans := 1
	for length := len(s) - 1; length > 1; length-- {
		mainCombs := genCombs(length, len(s))
		restLength := len(s) - length
		restCombs := make([][][]int, restLength)
		for subLength := restLength; subLength > 0; subLength-- {
			restCombs[subLength-1] = genCombs(subLength, restLength)
		}

		for subLength := restLength; subLength > 0; subLength-- {
			restCombs[subLength-1] = genCombs(subLength, restLength)
		}

		// when main is perm
	LOOP:
		for _, mainComb := range mainCombs {
			for x := 0; x < (len(mainComb) / 2); x++ {
				if s[mainComb[x]] != s[mainComb[len(mainComb)-1-x]] {
					continue LOOP
				}
			}

			m := map[int]bool{}
			for _, i := range mainComb {
				m[i] = true
			}

			restIndices := []int{}
			for i := 0; i < len(s); i++ {
				if m[i] {
					continue
				}
				restIndices = append(restIndices, i)
			}

			for _, restCombb := range restCombs {
			INNER:
				for _, restComb := range restCombb {
					for x := 0; x < (len(restComb) / 2); x++ {
						if s[restIndices[restComb[x]]] != s[restIndices[restComb[len(restComb)-1-x]]] {
							continue INNER
						}
					}
					if len(mainComb)*len(restComb) > ans {
						ans = len(mainComb) * len(restComb)
					}
				}
			}
		}
	}
	return ans
}

func genCombs(part, whole int) [][]int {
	ans := [][]int{}
	var perm func(arr []int, index int)
	perm = func(arr []int, index int) {
		if len(arr) == part {
			clone := make([]int, len(arr))
			copy(clone, arr)
			ans = append(ans, clone)
			return
		}
		if index == whole {
			return
		}
		perm(append(arr, index), index+1)
		perm(arr, index+1)
	}
	perm([]int{}, 0)
	return ans
}
