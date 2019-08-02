package hard

import (
	"fmt"
	"github.com/luci/go-render/render"
	"testing"
	"time"
)

/**
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

示例:

输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6
*/
func Test_mergeKLists(t *testing.T) {
	node10 := &ListNode{
		Val:1,
	}
	node11 := &ListNode{
		Val:4,
	}
	node12 := &ListNode{
		Val:5,
	}
	node10.Next = node11
	node11.Next = node12

	node20 := &ListNode{
		Val:1,
	}
	node21 := &ListNode{
		Val:3,
	}
	node22 := &ListNode{
		Val:4,
	}
	node20.Next = node21
	node21.Next = node22

	node30 := &ListNode{
		Val:2,
	}
	node31 := &ListNode{
		Val:6,
	}
	node30.Next = node31

	lists := make([]*ListNode,0)
	lists = append(lists,node10)
	lists = append(lists,node20)
	lists = append(lists,node30)
	got := mergeKLists(lists)
	fmt.Println(fmt.Sprintf("mergeKLists result is %v", got.Val))
	pointer := got
	for {
		if pointer.Next != nil {
			fmt.Println(fmt.Sprintf("%v", pointer.Val))
			pointer = pointer.Next
		} else {
			fmt.Println(fmt.Sprintf("%v", pointer.Val))
			break
		}
	}
	time.Sleep(1 * time.Second)

	lists1 := make([]*ListNode,2)
	got = mergeKLists(lists1)
	fmt.Println(fmt.Sprintf("mergeKLists result is %v", render.Render(got)))
	time.Sleep(1 * time.Second)

	lists2 := make([]*ListNode,2)
	node := &ListNode{
		Val:1,
	}
	lists2[1] = node
	got = mergeKLists(lists2)
	fmt.Println(fmt.Sprintf("mergeKLists result is %v", render.Render(got)))
	time.Sleep(1 * time.Second)

	lists3 := make([]*ListNode,3)
	node3 := &ListNode{
		Val:2,
	}
	node33 := &ListNode{
		Val:-1,
	}
	lists3[0] = node3
	lists3[2] = node33
	got = mergeKLists(lists3)
	fmt.Println(fmt.Sprintf("mergeKLists result is %v", render.Render(got)))
	time.Sleep(1 * time.Second)
}
