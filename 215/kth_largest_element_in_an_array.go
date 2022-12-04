// 215. Kth Largest Element in an Array
// https://leetcode.com/problems/kth-largest-element-in-an-array/
package main

import (
	"math/rand"
	"time"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	pp.Println(5)
	pp.Println("=========================================")
	pp.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
	pp.Println(4)
	pp.Println("=========================================")
	pp.Println(findKthLargest([]int{5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 3, 2, 1}, 8))
	pp.Println(5)
	pp.Println("=========================================")
}

/**
 * Constraints:
 * - 1 <= k <= nums.length <= 10^4
 * - -10^4 <= nums[i] <= 10^4
 *
 * K: Given value to determine the element to return
 * N: Length of the input array `nums`
 */

/**
 * 3. Shuffle the array, then Quick Select
 * - Time Complexity: O(N)
 * - Space Complexity: O(1)
 * > Runtime: 4 ms (faster than 98.32%)
 * > Memory Usage: 3.5 MB (less than 100.00%)
 */
func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(nums), func(i, j int) { nums[i], nums[j] = nums[j], nums[i] })
	for len(nums) > 1 {
		pivot := min(nums[0], nums[1])
		boundary := 0
		for i := range nums {
			if nums[i] > pivot {
				nums[i], nums[boundary] = nums[boundary], nums[i]
				boundary++
			}
		}
		if boundary == 0 {
			if k > 2 {
				boundary = 2
			} else {
				boundary = 1
			}
		}
		if k <= boundary {
			nums = nums[:boundary]
		} else {
			nums = nums[boundary:]
			k -= boundary
		}
	}
	return nums[0]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/**
 * 2. Use binary heap
 * - Time Complexity: O(NlogK)
 * - Space Complexity: O(1)
 * > Runtime: 0 ms (faster than 100.00%)
 * > Memory Usage: 3.5 MB (less than 100.00%)
 */
// type intHeap []int
//
// func (h intHeap) Len() int           { return len(h) }
// func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
// func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
//
// func (h *intHeap) Push(x interface{}) {
// 	*h = append(*h, x.(int))
// }
//
// func (h *intHeap) Pop() interface{} {
// 	old := *h
// 	n := len(old)
// 	x := old[n-1]
// 	*h = old[0 : n-1]
// 	return x
// }
//
// func findKthLargest(nums []int, k int) int {
// 	h := intHeap(nums[:k])
// 	heap.Init(&h)
// 	for _, num := range nums[k:] {
// 		if num > h[0] {
// 			h[0] = num
// 			heap.Fix(&h, 0)
// 		}
// 	}
// 	return h[0]
// }

/**
 * 1. Sort
 * - Time Complexity: O(NlogN)
 * - Space Complexity: O(logN)
 */
// func findKthLargest(nums []int, k int) int {
// 	sort.Ints(nums)
// 	return nums[len(nums)-k]
// }
