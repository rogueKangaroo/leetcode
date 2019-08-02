package medium
/**
在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

示例 1:

输入: [3,2,1,5,6,4] 和 k = 2
输出: 5
示例 2:

输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
输出: 4
说明:

你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。
解法：
	维持一个k容量的最小堆，只有比堆顶元素大的才能进堆
**/
func findKthLargest(nums []int, k int) int {
	heap := make([]int,k)
	for i:=0;i<k;i++ {
		heap[i] = nums[i]
	}
	for i:=k-1;i>=0;i-- {
		heapSort(&heap,k,i)
	}
	for i:=k;i<len(nums);i++ {
		if nums[i] > heap[0] {
			heap[0] = nums[i]
			heapSort(&heap,k,0)
		}
	}
	return heap[0]
}

func heapSort(heap *[]int,length int,index int)  {
	left := getLeftChildIndex(index)
	right := getRightChildIndex(index)
	min := index
	if left < length && (*heap)[left] < (*heap)[min] {
		min = left
	}
	if right < length && (*heap)[right] < (*heap)[min] {
		min = right
	}
	if min != index {
		(*heap)[min] = (*heap)[min] ^ (*heap)[index]
		(*heap)[index] = (*heap)[min] ^ (*heap)[index]
		(*heap)[min] = (*heap)[min] ^ (*heap)[index]
		heapSort(heap,length,min)
	}
}

func getLeftChildIndex(index int) int {
	return index << 1 + 1
}

func getRightChildIndex(index int) int {
	return index << 1 + 2
}