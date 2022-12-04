package helper

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func ListNodesToString(node *ListNode) string {
	var nums []string
	for n := node; n != nil; n = n.Next {
		nums = append(nums, strconv.Itoa(n.Val))
	}
	return fmt.Sprintf("[%s]", strings.Join(nums, ", "))
}
