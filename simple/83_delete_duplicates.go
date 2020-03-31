package simple

/*
给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

示例 1:

输入: 1->1->2
输出: 1->2
示例 2:

输入: 1->1->2->3->3
输出: 1->2->3

*/
type ListNode_deleteDuplicates struct {
     Val int
     Next *ListNode_deleteDuplicates
}

func deleteDuplicates(head *ListNode_deleteDuplicates) *ListNode_deleteDuplicates {
	if head == nil {
		return nil
	}
	pointer2 := head.Next
	pointer1 := head
	for {
		if pointer2 == nil {
			break
		}
		if pointer1.Val == pointer2.Val {
			pointer2 = pointer2.Next
			pointer1.Next = nil
		} else {
			pointer1.Next = pointer2
			pointer1 = pointer1.Next
			pointer2 = pointer2.Next
		}
	}
	return head
}
