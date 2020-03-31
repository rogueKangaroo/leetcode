package medium

import (
	"fmt"
	"strconv"
	"testing"
)

//给定 1->2->3->4, 你应该返回 2->1->4->3.
func Test_swapPairs(t *testing.T) {
	node3 := &ListNode_swapPairs{4,nil}
	node2 := &ListNode_swapPairs{3,node3}
	node1 := &ListNode_swapPairs{2,node2}
	head := &ListNode_swapPairs{1,node1}
	got := swapPairs(head)
	for {
		if got == nil {
			break
		} else {
			fmt.Print(strconv.Itoa(got.Val)+"->")
			got = got.Next
		}
	}
}
