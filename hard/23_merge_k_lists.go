package hard

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

解法：
	建造一个K容量的最小堆，堆顶node，推出（最小元素），
	如果堆顶node有next结点，则堆顶元素变为next结点，然后构造最小堆，
	如果没有最小结点，则跟堆末元素交换位置推出，然后构造最小堆
**/

type ListNode struct {
     Val int
     Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	newLists := handleLists(lists)
	if len(newLists) == 0 {
		return nil
	}
	for i:=len(newLists) - 1;i >= 0;i-- {
		heapSort(&newLists,len(newLists),i)
	}
	header := newLists[0]
	pointer := header
	for i:=len(newLists) - 1;i > 0;i-- {
		if newLists[0] != nil {
			if newLists[0].Next != nil {
				newLists[0] = newLists[0].Next
				i ++
				heapSort(&newLists,i,0)
				pointer.Next = newLists[0]
				pointer = pointer.Next
			} else {
				temp := newLists[0]
				newLists[0] = newLists[i]
				newLists[i] = temp
				heapSort(&newLists,i,0)
				pointer.Next = newLists[0]
				pointer = pointer.Next
			}
		}
	}
	return header
}

func heapSort(lists *[]*ListNode,length int, index int)  {
	left := getLeftChild(index)
	right := getRightChild(index)
	min := index
	if(left < length && (*lists)[left] != nil && (*lists)[index] != nil && (*lists)[left].Val < (*lists)[min].Val){
		min=left
	}
	if(right < length && (*lists)[right] != nil && (*lists)[index] != nil && (*lists)[right].Val < (*lists)[min].Val){
		min=right
	}
	if(min != index){
		temp := (*lists)[index]
		(*lists)[index] = (*lists)[min]
		(*lists)[min] = temp
		heapSort(lists,length,min)
	}
}

func getRightChild(index int) int {
	return (index<<1)+2
}

func getLeftChild(index int) int {
	return (index<<1)+1
}

func handleLists(lists []*ListNode) []*ListNode {
	newList := make([]*ListNode,0)
	for i:=0;i<len(lists);i++ {
		if lists[i] != nil {
			newList = append(newList,lists[i])
		}
	}
	return newList
}