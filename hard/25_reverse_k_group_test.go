package hard

import (
	"fmt"
	"strconv"
	"testing"
)

//给你这个链表：1->2->3->4->5
//
//当 k = 2 时，应当返回: 2->1->4->3->5
//
//当 k = 3 时，应当返回: 3->2->1->4->5
func Test_reverseKGroup(t *testing.T) {
	node5 := &ListNode_reverseKGroup{5,nil}
	node4 := &ListNode_reverseKGroup{4,node5}
	node3 := &ListNode_reverseKGroup{3,node4}
	node2 := &ListNode_reverseKGroup{2,node3}
	head := &ListNode_reverseKGroup{1,node2}
	k := 2
	got := reverseKGroup(head,k)
	for {
		if got != nil {
			fmt.Print(strconv.Itoa(got.Val)+"->")
			got = got.Next
		} else {
			break
		}
	}
}


//给你这个链表：1->2
//
//当 k = 2 时，应当返回: 2->1
func Test_reverseKGroup1(t *testing.T) {
	node2 := &ListNode_reverseKGroup{2,nil}
	head := &ListNode_reverseKGroup{1,node2}
	k := 2
	got := reverseKGroup(head,k)
	for {
		if got != nil {
			fmt.Print(strconv.Itoa(got.Val)+"->")
			got = got.Next
		} else {
			break
		}
	}
}