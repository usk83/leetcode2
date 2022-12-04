// 5839. Remove Stones to Minimize the Total
// https://leetcode.com/contest/weekly-contest-253/problems/remove-stones-to-minimize-the-total/
package main

import (
	"container/heap"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(minStoneSum([]int{5, 4, 9}, 2))
	pp.Println(12)
	pp.Println("=========================================")
	pp.Println(minStoneSum([]int{4, 3, 6, 7}, 3))
	pp.Println(12)
	pp.Println("=========================================")
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minStoneSum(piles []int, k int) int {
	sum := 0
	for _, pile := range piles {
		sum += pile
	}

	ih := IntHeap(piles)
	h := &ih
	heap.Init(h)

	for i := 0; i < k; i++ {
		max := heap.Pop(h).(int)
		half := max / 2
		sum -= half
		heap.Push(h, max-half)
	}

	return sum
}
