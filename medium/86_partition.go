package medium

/*
给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。

你应当保留两个分区中每个节点的初始相对位置。

示例:

输入: head = 1->4->3->2->5->2, x = 3
输出: 1->2->2->4->3->5
解法：分隔成两个链表，即大于等于一个链表，小于一个链表，然后再合并
*/


type ListNode_partition struct {
     Val int
     Next *ListNode_partition
}

func partition(head *ListNode_partition, x int) *ListNode_partition {
	if head == nil {
		return nil
	}
	maxNodeList := make([]*ListNode_partition,0)
	minNodeList := make([]*ListNode_partition,0)
	node := head
	for {
		if node == nil {
			break
		} else if node.Val >= x {
			maxNodeList = append(maxNodeList,node)
		} else {
			minNodeList = append(minNodeList,node)
		}
		node = node.Next
	}
	resultNode := &ListNode_partition{}
	resultHead := &ListNode_partition{}
	headFlag := true
	for i:=0;i<len(minNodeList);i++ {
		if i == 0 {
			resultNode = minNodeList[0]
			resultHead = resultNode
			headFlag = false
		} else {
			resultNode.Next = minNodeList[i]
			resultNode = resultNode.Next
		}
	}

	for i:=0;i<len(maxNodeList);i++ {
		if headFlag && i==0{
			resultNode = maxNodeList[0]
			resultHead = resultNode
			headFlag = false
		} else {
			resultNode.Next = maxNodeList[i]
			resultNode = resultNode.Next
		}
		if i == len(maxNodeList) - 1 {
			resultNode.Next = nil
		}
	}
	return resultHead
}