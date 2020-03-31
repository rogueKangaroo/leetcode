package simple

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

//输入: 1->1->2
//输出: 1->2
func Test_deleteDuplicates(t *testing.T) {
	node2 := &ListNode_deleteDuplicates{2,nil}
	node1 := &ListNode_deleteDuplicates{1,node2}
	head := &ListNode_deleteDuplicates{1,node1}
	got := deleteDuplicates(head)
	for {
		if got == nil {
			break
		} else {
			fmt.Print(strconv.Itoa(got.Val)+"->")
			got = got.Next
		}
	}
	time.Sleep(1*time.Second)
}

//输入: 1->1->2->3->3
//输出: 1->2->3
func Test_deleteDuplicates1(t *testing.T) {
	node4 := &ListNode_deleteDuplicates{3,nil}
	node3 := &ListNode_deleteDuplicates{3,node4}
	node2 := &ListNode_deleteDuplicates{2,node3}
	node1 := &ListNode_deleteDuplicates{1,node2}
	head := &ListNode_deleteDuplicates{1,node1}
	got := deleteDuplicates(head)
	for {
		if got == nil {
			break
		} else {
			fmt.Print(strconv.Itoa(got.Val)+"->")
			got = got.Next
		}
	}
	time.Sleep(1*time.Second)
}