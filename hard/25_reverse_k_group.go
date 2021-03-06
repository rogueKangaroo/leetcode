package hard


/*

给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。

k 是一个正整数，它的值小于或等于链表的长度。

如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

 

示例：

给你这个链表：1->2->3->4->5

当 k = 2 时，应当返回: 2->1->4->3->5

当 k = 3 时，应当返回: 3->2->1->4->5

 

说明：

你的算法只能使用常数的额外空间。
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
*/

type ListNode_reverseKGroup struct {
     Val int
     Next *ListNode_reverseKGroup
}

func reverseKGroup(head *ListNode_reverseKGroup, k int) *ListNode_reverseKGroup {
	nodeList := make([]*ListNode_reverseKGroup,k)
	isHead := true
	pointerNode := head
	flagNode := head
	for {
		for i:=0;i<k;i++ {
			if pointerNode == nil && i<k{
				return head
			}
			nodeList[i] = pointerNode
			pointerNode = pointerNode.Next
		}
		flagNode.Next = nodeList[k-1]
		if isHead {
			head = nodeList[k-1]
			isHead = false
		}
		for i:=k-1;i>=1;i--{
			nodeList[i].Next = nodeList[i-1]
		}
		nodeList[0].Next = pointerNode
		flagNode = nodeList[0]
	}

	return head
}