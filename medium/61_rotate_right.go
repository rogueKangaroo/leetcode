package medium

/*
给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。

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

解法：先找到listNode的长度，取余得到移动几步，然后找到最后的k节点，进行断接
*/

type ListNode_rotateRight struct {
     Val int
     Next *ListNode_rotateRight
}

func rotateRight(head *ListNode_rotateRight, k int) *ListNode_rotateRight {
	if head == nil || k == 0{
		return head
	}
	len := 1
	node := head
	for {
		if node.Next != nil {
			len ++
			node = node.Next
		} else {
			break
		}
	}
	realK := k%len
	if realK == 0 {
		return head
	}
	flagK := head
	tailNode := flagK
	nodeK := head
	for {
		if nodeK.Next != nil {
			nodeK = nodeK.Next
			realK --
			if realK <=0 {
				tailNode = flagK
				flagK = flagK.Next
			}
		} else {
			break
		}
	}
	nodeK.Next = head
	tailNode.Next = nil
	return flagK
}
