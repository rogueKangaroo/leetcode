package medium

import (
	"fmt"
	"github.com/luci/go-render/render"
	"testing"
	"time"
)

/*
示例 1:

输入: 1->2->3->4->5->NULL, k = 2
输出: 4->5->1->2->3->NULL
解释:
向右旋转 1 步: 5->1->2->3->4->NULL
向右旋转 2 步: 4->5->1->2->3->NULL
示例 2:

输入: 0->1->2->NULL, k = 4
输出: 2->0->1->NULL
解释:
向右旋转 1 步: 2->0->1->NULL
向右旋转 2 步: 1->2->0->NULL
向右旋转 3 步: 0->1->2->NULL
向右旋转 4 步: 2->0->1->NULL

*/
func Test_rotateRight(t *testing.T) {
	node4 := &ListNode_rotateRight{Val:5,Next:nil}
	node3 := &ListNode_rotateRight{Val:4,Next:node4}
	node2 := &ListNode_rotateRight{Val:3,Next:node3}
	node1 := &ListNode_rotateRight{Val:2,Next:node2}
	head := &ListNode_rotateRight{Val:1,Next:node1}
	k := 2
	got := rotateRight(head,k)
	fmt.Print(fmt.Sprintf("list is "))
	for{
		if got.Next != nil{
			fmt.Print(fmt.Sprintf("->%v", render.Render(got.Val)))
			got = got.Next
		}else {
			fmt.Print(fmt.Sprintf("->%v", render.Render(got.Val)))
			break
		}
	}
	time.Sleep(1*time.Second)
}

func Test_rotateRight1(t *testing.T) {
	node2 := &ListNode_rotateRight{Val:2,Next:nil}
	node1 := &ListNode_rotateRight{Val:1,Next:node2}
	head := &ListNode_rotateRight{Val:0,Next:node1}
	k := 4
	got := rotateRight(head,k)
	fmt.Print(fmt.Sprintf("list is "))
	for{
		if got.Next != nil{
			fmt.Print(fmt.Sprintf("->%v", render.Render(got.Val)))
			got = got.Next
		}else {
			fmt.Print(fmt.Sprintf("->%v", render.Render(got.Val)))
			break
		}
	}
	time.Sleep(1*time.Second)
}

func Test_rotateRight2(t *testing.T) {
	head := &ListNode_rotateRight{Val:0,Next:nil}
	k := 0
	got := rotateRight(head,k)
	fmt.Print(fmt.Sprintf("list is "))
	for{
		if got.Next != nil{
			fmt.Print(fmt.Sprintf("->%v", render.Render(got.Val)))
			got = got.Next
		}else {
			fmt.Print(fmt.Sprintf("->%v", render.Render(got.Val)))
			break
		}
	}
	time.Sleep(1*time.Second)
}

func Test_rotateRight3(t *testing.T) {
	node := &ListNode_rotateRight{Val:2,Next:nil}
	head := &ListNode_rotateRight{Val:1,Next:node}
	k := 2
	got := rotateRight(head,k)
	fmt.Print(fmt.Sprintf("list is "))
	for{
		if got.Next != nil{
			fmt.Print(fmt.Sprintf("->%v", render.Render(got.Val)))
			got = got.Next
		}else {
			fmt.Print(fmt.Sprintf("->%v", render.Render(got.Val)))
			break
		}
	}
	time.Sleep(1*time.Second)
}