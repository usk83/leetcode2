// 49. Group Anagrams
// https://leetcode.com/problems/group-anagrams/
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println([][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}})
	pp.Println("=========================================")
	fmt.Println(groupAnagrams([]string{""}))
	fmt.Println([][]string{{""}})
	pp.Println("=========================================")
	fmt.Println(groupAnagrams([]string{"a"}))
	fmt.Println([][]string{{"a"}})
	pp.Println("=========================================")
}

/**
 * Constraints:
 * - 1 <= strs.length <= 10^4
 * - 0 <= strs[i].length <= 100
 * - strs[i] consists of lowercase English letters.
 *
 * N: Length of the input array `strs`
 * M: Total number of characters of given strings
 */

/**
 * 2. Group directly into a slice while managing the correspondence between keys and indexes
 * - Time Complexity: O(M)
 * - Space Complexity: O(N)
 * > Runtime: 20 ms (faster than 96.36%)
 * > Memory Usage: 7.3 MB (less than 97.64%)
 */
func groupAnagrams(strs []string) [][]string {
	groupedStrs := [][]string{}
	groupIDs := map[[26]uint8]uint16{}
	for _, str := range strs {
		key := [26]uint8{}
		for _, r := range str {
			key[r-'a']++
		}
		if _, ok := groupIDs[key]; !ok {
			groupIDs[key] = uint16(len(groupedStrs))
			groupedStrs = append(groupedStrs, nil)
		}
		groupedStrs[groupIDs[key]] = append(groupedStrs[groupIDs[key]], str)
	}
	return groupedStrs
}

/**
 * 1. Group by the count of each letters in each strings using a hashmap
 *    and convert it to a slice
 * - Time Complexity: O(M)
 * - Space Complexity: O(N)
 * > Runtime: 20 ms (faster than 96.36%)
 * > Memory Usage: 8.6 MB (less than 47.09%)
 */
// func groupAnagrams(strs []string) [][]string {
// 	groupedMap := map[[26]int][]string{}
// 	for _, str := range strs {
// 		key := [26]int{}
// 		for _, r := range str {
// 			key[r-'a']++
// 		}
// 		groupedMap[key] = append(groupedMap[key], str)
// 	}
// 	groupedSlice := make([][]string, 0, len(groupedMap))
// 	for _, v := range groupedMap {
// 		groupedSlice = append(groupedSlice, v)
// 	}
// 	return groupedSlice
// }
