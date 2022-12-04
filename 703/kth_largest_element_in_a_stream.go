// 703. Kth Largest Element in a Stream
// https://leetcode.com/problems/kth-largest-element-in-a-stream/
package main

import (
	"container/heap"

	"github.com/k0kubun/pp"
)

func main() {
	kth := Constructor(3, []int{4, 5, 8, 2})
	pp.Println("=========================================")
	pp.Println(kth.Add(3))
	pp.Println(4)
	pp.Println("=========================================")
	pp.Println(kth.Add(5))
	pp.Println(5)
	pp.Println("=========================================")
	pp.Println(kth.Add(10))
	pp.Println(5)
	pp.Println("=========================================")
	pp.Println(kth.Add(9))
	pp.Println(8)
	pp.Println("=========================================")
	pp.Println(kth.Add(4))
	pp.Println(8)
	pp.Println("=========================================")
}

/**
 * Constraints:
 * - 1 <= k <= 10^4
 * - 0 <= nums.length <= 10^4
 * - -10^4 <= nums[i] <= 10^4
 * - -10^4 <= val <= 10^4
 * - At most 10^4 calls will be made to add.
 * - It is guaranteed that there will be at least k elements in the array when you search for the kth element.
 *
 * K: Value for finding an element
 * N: Length of the input array `nums`
 * M: Number of calls to `add`
 */

/**
 * 4. Use binary heap to keep elements
 * - Time Complexity  (init): O(NlogN)
 * - Time Complexity (calls): O(MlogK)
 * - Space Complexity: O(K)
 * > Runtime: 28 ms (faster than 97.30%)
 * > Memory Usage: 8 MB (less than 85.13%)
 */
const (
	MaxInt = int(^uint(0) >> 1)
	MinInt = -MaxInt - 1
)

/* https://pkg.go.dev/container/heap */
type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type KthLargest struct {
	elements intHeap
}

func Constructor(k int, nums []int) KthLargest {
	elements := intHeap(nums)
	heap.Init(&elements)

	if len(nums) < k {
		heap.Push(&elements, MinInt)
	} else {
		for i := len(nums); i > k; i-- {
			_ = heap.Pop(&elements)
		}
	}

	return KthLargest{
		elements: elements,
	}
}

func (this *KthLargest) Add(val int) int {
	if val > this.elements[0] {
		heap.Push(&this.elements, val)
		_ = heap.Pop(&this.elements)
	}

	return this.elements[0]
}

/**
 * 3. Keep `k` elements and do insertion sort every time on `add`
 * - Time Complexity  (init): O(NlogN)
 * - Time Complexity (calls): O(MK)
 * - Space Complexity: O(K) (or O(logN) because of sorting on init)
 */
// const (
// 	MaxInt = int(^uint(0) >> 1)
// 	MinInt = -MaxInt - 1
// )
//
// type KthLargest struct {
// 	elements []int
// }
//
// func Constructor(k int, nums []int) KthLargest {
// 	if k > len(nums) {
// 		nums = append(nums, MinInt)
// 	}
// 	sort.Ints(nums)
// 	return KthLargest{
// 		elements: nums[len(nums)-k : len(nums)],
// 	}
// }
//
// func (this *KthLargest) Add(val int) int {
// 	if val > this.elements[0] {
// 		target := sort.Search(len(this.elements), func(i int) bool {
// 			return this.elements[i] >= val
// 		})
// 		for i := target - 1; i >= 0; i-- {
// 			this.elements[i], val = val, this.elements[i]
// 		}
// 	}
// 	return this.elements[0]
// }

/**
 * 2. Sort every time on `add`, but keep only `k` elements
 * - Time Complexity  (init): O(NlogN)
 * - Time Complexity (calls): O(MKlogK)
 * - Space Complexity: O(K) (or O(min(logN, logK)) because of sorting)
 */
// const (
// 	MaxInt = int(^uint(0) >> 1)
// 	MinInt = -MaxInt - 1
// )
//
// type KthLargest struct {
// 	elements []int
// }
//
// func Constructor(k int, nums []int) KthLargest {
// 	if k > len(nums) {
// 		nums = append(nums, MinInt)
// 	}
// 	sort.Ints(nums)
// 	return KthLargest{
// 		elements: nums[len(nums)-k : len(nums)],
// 	}
// }
//
// func (this *KthLargest) Add(val int) int {
// 	if val > this.elements[0] {
// 		this.elements[0] = val
// 		sort.Ints(this.elements)
// 	}
// 	return this.elements[0]
// }

/**
 * 1. Sort every time on `add`
 * - Time Complexity: O(Mlog(N+M))
 * - Space Complexity: O(N+M)
 */
// type KthLargest struct {
// 	k        int
// 	elements []int
// }
//
// func Constructor(k int, nums []int) KthLargest {
// 	return KthLargest{
// 		k:        k,
// 		elements: nums,
// 	}
// }
//
// func (this *KthLargest) Add(val int) int {
// 	this.elements = append(this.elements, val)
// 	sort.Ints(this.elements)
// 	return this.elements[len(this.elements)-this.k]
// }
