package medium

/*
假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,0,1,2,2,5,6] 可能变为 [2,5,6,0,0,1,2] )。

编写一个函数来判断给定的目标值是否存在于数组中。若存在返回 true，否则返回 false。

示例 1:

输入: nums = [2,5,6,0,0,1,2], target = 0
输出: true
示例 2:

输入: nums = [2,5,6,0,0,1,2], target = 3
输出: false
进阶:

这是 搜索旋转排序数组 的延伸题目，本题中的 nums  可能包含重复元素。
这会影响到程序的时间复杂度吗？会有怎样的影响，为什么？

解法：二分查找
*/
func search81(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	if nums[0] < nums[len(nums) -1] {
		return searchBinary(nums,target,0,len(nums) - 1)
	} else {
		return searchBinaryTransfer(nums,target,0,len(nums) - 1)
	}
}

func searchBinaryTransfer(nums []int, target int, start,end int) bool {
	index :=  (start + end)/2
	if index == start {
		if nums[index] == target || nums[end] == target {
			return true
		} else {
			return false
		}
	}
	if nums[index] == target {
		return true
	} else if nums[index] < target {
		if nums[index] > nums[len(nums)-1] {
			return searchBinaryTransfer(nums,target,index,end)
		} else if nums[index] == nums[len(nums)-1] {
			return searchBinaryTransfer(nums,target,start,index) || searchBinaryTransfer(nums,target,index,end)
		} else {
			return searchBinaryTransfer(nums,target,start,index) || searchBinary(nums,target,index,end)
		}
	} else {
		if nums[index] < nums[0] {
			return searchBinaryTransfer(nums,target,start,index)
		} else if nums[index] < nums[0] {
			return searchBinaryTransfer(nums,target,start,index) || searchBinaryTransfer(nums,target,index,end)
		} else {
			return searchBinaryTransfer(nums,target,index,end) || searchBinary(nums,target,start,end)
		}
	}
}

func searchBinary(nums []int, target int, start,end int) bool {
	index :=  (start + end)/2
	if index == start {
		if nums[index] == target || nums[end] == target {
			return true
		} else {
			return false
		}
	}
	if nums[index] == target {
		return true
	} else if nums[index] < target {
		return searchBinary(nums,target,index,end)
	} else {
		return searchBinary(nums,target,start,index)
	}
}
