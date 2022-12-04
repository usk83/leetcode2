// 5851. Find Unique Binary String
// https://leetcode.com/contest/weekly-contest-255/problems/find-unique-binary-string/
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(findDifferentBinaryString([]string{"10", "01"}))
	pp.Println()
	pp.Println("=========================================")
	pp.Println(findDifferentBinaryString([]string{"00", "11"}))
	pp.Println()
	pp.Println("=========================================")
	pp.Println(findDifferentBinaryString([]string{"111", "000", "001"}))
	pp.Println()
	pp.Println("=========================================")
	pp.Println(findDifferentBinaryString([]string{"0000", "0011", "0001", "0010"}))
	pp.Println()
	pp.Println("=========================================")
}

/**
 * [description]
 * - Time Complexity: O(X)
 * - Space Complexity: O(X)
 */
func findDifferentBinaryString(nums []string) string {
	n := len(nums)
	pow := 1 << uint(n)
LOOP:
	for i := 0; i < pow; i++ {
		bin := fmt.Sprintf("%b", i)
		for i := n - len(bin); i > 0; i-- {
			bin = "0" + bin
		}
		for _, num := range nums {
			if num == bin {
				continue LOOP
			}
		}
		return bin
	}
	return "invalid"
}
