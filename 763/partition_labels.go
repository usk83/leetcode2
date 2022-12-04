// 763. Partition Labels
// https://leetcode.com/problems/partition-labels/
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
	fmt.Println([]int{9, 7, 8})
	pp.Println("=========================================")
	fmt.Println(partitionLabels("eccbbbbdec"))
	fmt.Println([]int{10})
	pp.Println("=========================================")
}

/**
 * 1. Calculate the last appearance of each characters, then partition the input string using it
 * - Time Complexity: O(N)
 * - Space Complexity: O(1)
 * > Runtime: 0 ms (faster than 100.00%)
 * > Memory Usage: 2.2 MB (less than 100.00%)
 */
func partitionLabels(s string) []int {
	var lastAppearance [26]int
	for i, b := range s {
		lastAppearance[b-'a'] = i
	}

	partition, sum, lastIndex := []int{}, 0, -1
	for i, b := range s {
		lastIndex = max(lastIndex, lastAppearance[b-'a'])
		if i == lastIndex {
			size := i + 1 - sum
			partition = append(partition, size)
			sum += size
		}
	}
	return partition
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/**
 * 99. Count the number of each characters' appearance, then partition the input string using it
 *     (A written code of my friend's idea)
 * - Time Complexity: O(N) since the number of character types is fixed at 26
 * - Space Complexity: O(1) for the same reason as above
 * > Runtime: 12 ms (faster than 9.09%)
 * > Memory Usage: 2.5 MB (less than 18.18%)
 */
// func partitionLabels(s string) []int {
// 	var count [26]int
// 	for _, b := range s {
// 		count[b-'a']++
// 	}
// 	partition, sum, currentCountMap := []int{}, 0, map[rune]int{}
// LOOP:
// 	for i, b := range s {
// 		currentCountMap[b]++
// 		for k, v := range currentCountMap {
// 			if v != count[k-'a'] {
// 				continue LOOP
// 			}
// 		}
// 		size := i + 1 - sum
// 		partition = append(partition, size)
// 		sum += size
// 		currentCountMap = map[rune]int{}
// 	}
// 	return partition
// }
