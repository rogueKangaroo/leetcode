package medium

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

//给定一个链表: 1->2->3->4->5, 和 n = 2.
func Test_removeNthFromEnd(t *testing.T) {
	node5 := &ListNode_removeNthFromEnd{5,nil}
	node4 := &ListNode_removeNthFromEnd{4,node5}
	node3 := &ListNode_removeNthFromEnd{3,node4}
	node2 := &ListNode_removeNthFromEnd{2,node3}
	head := &ListNode_removeNthFromEnd{1,node2}
	n := 2
	got := removeNthFromEnd(head, n)
	for {
		if got == nil {
			break
		} else {
			fmt.Print(strconv.Itoa(got.Val)+"->")
			got = got.Next
		}
	}
	time.Sleep(1 * time.Second)
}


//给定一个链表: 1->2->3->4->5, 和 n = 2.
func Test_removeNthFromEnd2(t *testing.T) {
	node2 := &ListNode_removeNthFromEnd{2,nil}
	head := &ListNode_removeNthFromEnd{1,node2}
	n := 2
	got := removeNthFromEnd(head, n)
	for {
		if got == nil {
			break
		} else {
			fmt.Print(strconv.Itoa(got.Val)+"->")
			got = got.Next
		}
	}
	time.Sleep(1 * time.Second)
}

//给定一个链表: 1->2->3->4->5, 和 n = 2.
func Test_removeNthFromEnd1(t *testing.T) {
	head := &ListNode_removeNthFromEnd{1,nil}
	n := 1
	got := removeNthFromEnd(head, n)
	for {
		if got == nil {
			break
		} else {
			fmt.Print(strconv.Itoa(got.Val)+"->")
			got = got.Next
		}
	}
	time.Sleep(1 * time.Second)
}