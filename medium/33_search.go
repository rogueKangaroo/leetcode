package medium
/*
假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。

你可以假设数组中不存在重复的元素。

你的算法时间复杂度必须是 O(log n) 级别。

示例 1:

输入: nums = [4,5,6,7,0,1,2], target = 0
输出: 4
示例 2:

输入: nums = [4,5,6,7,0,1,2], target = 3
输出: -1

解法：由复杂度推断，必须采用二分法查找，
242
*/
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	} else if len(nums) == 1 {
		if nums[0] == target{
			return 0
		} else {
			return -1
		}
	}
	return searchFromNums(nums,0,len(nums)-1,target)
}

func searchFromNums(nums []int,start int,end int,target int) int{
	if start == end {
		return -1
	}
	index := (start+end)/2
	if nums[index] > nums[len(nums)-1] {
		// 左边是有序的
		if nums[index] == target {
			return index
		} else {
			sortResult :=  searchSorted(nums,start,index,target)
			noSortedResult := -1
			if index == start {
				if nums[end] == target {
					return end
				} else {
					return -1
				}
			} else {
				noSortedResult = searchFromNums(nums,index,end,target)
			}
			if sortResult == -1 && noSortedResult == -1 {
				return -1
			} else if sortResult > noSortedResult {
				return sortResult
			} else {
				return noSortedResult
			}
		}
	} else if nums[index] < nums[len(nums) - 1] {
		// 右边是有序的
		if nums[index] == target {
			return index
		} else {
			sortResult :=  searchSorted(nums,index,end,target)
			noSortedResult := searchFromNums(nums,start,index,target)
			if sortResult == -1 && noSortedResult == -1 {
				return -1
			} else if sortResult > noSortedResult {
				return sortResult
			} else {
				return noSortedResult
			}
		}
	} else {
		if nums[index] == target {
			return index
		} else {
			return -1
		}
	}
}

func searchSorted(nums []int,start int,end int,target int) int {
	if start == end {
		return -1
	}
	index := (start+end)/2
	if nums[index] > target {
		return searchSorted(nums,start,index,target)
	} else if nums[index] < target {
		if index == start {
			if nums[end] == target {
				return end
			} else {
				return -1
			}
		} else {
			return searchSorted(nums,index,end,target)
		}
	} else {
		return index
	}
}
