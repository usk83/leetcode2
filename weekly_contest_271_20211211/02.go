// 5953. Sum of Subarray Ranges
// https://leetcode.com/contest/weekly-contest-271/problems/sum-of-subarray-ranges/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(subArrayRanges([]int{1, 2, 3}))
	pp.Println(4)
	pp.Println("=========================================")
	pp.Println(subArrayRanges([]int{1, 3, 3}))
	pp.Println(4)
	pp.Println("=========================================")
	pp.Println(subArrayRanges([]int{4, -2, -3, 4, 1}))
	pp.Println(59)
	pp.Println("=========================================")
}

// 1 <= nums.length <= 1000
// -10^9 <= nums[i] <= 10^9

// type num struct {
// 	value int
// 	index int
// }

// type minHeap []*num

// func (h minHeap) Len() int            { return len(h) }
// func (h minHeap) Less(i, j int) bool  { return h[i].value < h[j].value }
// func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
// func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(*num)) }
// func (h *minHeap) Pop() interface{} {
// 	old := *h
// 	length := len(old)
// 	n := old[length-1]
// 	old[length-1] = nil // avoid memory leak
// 	*h = old[0 : length-1]
// 	return n
// }

// type maxHeap []*num

// func (h minHeap) Len() int            { return len(h) }
// func (h minHeap) Less(i, j int) bool  { return h[i].value > h[j].value }
// func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
// func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(*num)) }
// func (h *minHeap) Pop() interface{} {
// 	old := *h
// 	length := len(old)
// 	n := old[length-1]
// 	old[length-1] = nil // avoid memory leak
// 	*h = old[0 : length-1]
// 	return n
// }

// func subArrayRanges(nums []int) int64 {
// 	var sum int64
// 	for start := 0; start < len(nums)-1; start++ {
// 		maxh, minh := maxHeap{nums[start]}, minHeap{nums[start]}
// 		for end := start+1; end < len(nums); end++ {

// 		}
// 	}
// 	return sum
// }

func subArrayRanges(nums []int) int64 {
	var sum int64
	for start := 0; start < len(nums)-1; start++ {
		minNum, maxNum := nums[start], nums[start]
		for end := start + 1; end < len(nums); end++ {
			minNum, maxNum := min(minNum, nums[end]), max(maxNum, nums[end])
			sum += int64(maxNum - minNum)
		}
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
