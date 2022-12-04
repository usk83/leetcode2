// 1. Two Sum
// https://leetcode.com/problems/two-sum/
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println([]int{0, 1})
	pp.Println("=========================================")
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println([]int{1, 2})
	pp.Println("=========================================")
	fmt.Println(twoSum([]int{3, 3}, 6))
	fmt.Println([]int{0, 1})
	pp.Println("=========================================")
	fmt.Println(twoSum([]int{2, 5, 5, 11}, 10))
	fmt.Println([]int{1, 2})
	pp.Println("=========================================")
	fmt.Println(twoSum([]int{-1, -2, -3, -4, -5}, -8))
	fmt.Println([]int{2, 4})
	pp.Println("=========================================")
}

/**
 * 4. One-pass solution using HashMap
 * - Time Complexity: O(N)
 * - Space Complexity: O(N)
 * > Runtime: 4 ms (faster than 97.45%)
 * > Memory Usage: 4.2 MB (less than 62.99%)
 */
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums)) // or `m := map[int]int{}`
	for i, num := range nums {
		if j, ok := m[target-num]; ok {
			return []int{i, j}
		}
		m[num] = i
	}
	return nil // never reached
}

/**
 * 3. Two-pass solution using HashMap
 * - Time Complexity: O(N)
 * - Space Complexity: O(N)
 */
// func twoSum(nums []int, target int) []int {
// 	m := map[int][]int{}
// 	for i, num := range nums {
// 		m[num] = append(m[num], i)
// 	}
// 	for i, num := range nums {
// 		if pair := target - num; num == pair {
// 			if len(m[num]) == 2 {
// 				return m[num]
// 			}
// 		} else {
// 			if s, ok := m[pair]; ok {
// 				return []int{i, s[0]}
// 			}
// 		}
// 	}
// 	return nil
// }

/**
 * 2. Sort, then binary search
 * - Time Complexity: O(NlogN)
 * - Space Complexity: O(N)
 */
// type elem struct {
// 	num   int
// 	index int
// }
//
// func twoSum(nums []int, target int) []int {
// 	arr := make([]elem, len(nums))
// 	for index, num := range nums {
// 		arr[index] = elem{num, index}
// 	}
// 	sort.Slice(arr, func(i, j int) bool { return arr[i].num < arr[j].num })
// 	for l, r := 0, len(nums)-1; l < r; {
// 		if sum := arr[l].num + arr[r].num; sum == target {
// 			return []int{arr[l].index, arr[r].index}
// 		} else if sum < target {
// 			l += 1 + sort.Search(r-l-2, func(i int) bool { return arr[l+1+i].num+arr[r].num >= target })
// 		} else {
// 			r -= 1 + sort.Search(r-l-2, func(i int) bool { return arr[l].num+arr[r-1-i].num <= target })
// 		}
// 	}
// 	return nil
// }

/**
 * 1. Naive solution (double loop)
 * - Time Complexity: O(N^2)
 * - Space Complexity: O(1)
 */
// func twoSum(nums []int, target int) []int {
// 	for i := range nums {
// 		for j := range nums[i+1:] {
// 			if k := i + 1 + j; nums[i]+nums[k] == target {
// 				return []int{i, k}
// 			}
// 		}
// 	}
// 	return nil
// }
