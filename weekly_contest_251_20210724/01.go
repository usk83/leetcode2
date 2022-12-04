// 5823. Sum of Digits of String After Convert
// https://leetcode.com/contest/weekly-contest-251/problems/sum-of-digits-of-string-after-convert/
package main

import (
	"strconv"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(getLucky("iiii", 1))
	pp.Println(36)
	pp.Println("=========================================")
	pp.Println(getLucky("leetcode", 2))
	pp.Println(6)
	pp.Println("=========================================")
	pp.Println(getLucky("zbax", 2))
	pp.Println(8)
	pp.Println("=========================================")
	pp.Println(getLucky("fleyctuuajsr", 5))
	pp.Println(8)
	pp.Println("=========================================")
}

func getLucky(s string, k int) int {
	var org string
	for _, b := range s {
		org += strconv.Itoa(int(b) - 96)
	}

	cur := org
	for i := 0; i < k; i++ {
		var sum int
		for _, b := range cur {
			sum += int(b) - 48
		}

		cur = strconv.Itoa(sum)
	}

	ans, _ := strconv.Atoi(cur)
	return ans
}
