package medium

/*
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

你的算法时间复杂度必须是 O(log n) 级别。

如果数组中不存在目标值，返回 [-1, -1]。

示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]
示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]

解法：根据复杂度来看必须采用二分，先二分查找找到target,然后前后继续二分，此时二分要分前二分（优先在前部分找到target）和后二分(优先在后部分找到target)
948
*/
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1,-1}
	} else if len(nums) == 1 {
		if nums[0] == target {
			return []int{0,0}
		}else {
			return []int{-1,-1}
		}
	}
	flagIndex := searchNums(nums,0,len(nums)-1,target)
	if flagIndex == -1{
		return []int{-1,-1}
	} else {
		start := searchBefore(nums,0,flagIndex,target)
		end := searchAfter(nums,flagIndex,len(nums)-1,target)
		if start == -1 && end == -1{
			return []int{flagIndex,flagIndex}
		} else if start != -1 && end == -1{
			return []int{start,flagIndex}
		} else if start == -1 && end != -1 {
			return []int{flagIndex,end}
		} else {
			return []int{start,end}
		}
	}
}

func searchNums(nums []int,start,end,target int) int {
	if start == end {
		return -1
	}
	index := (start + end)/2
	if nums[index] > target {
		return searchNums(nums,start,index,target)
	} else if nums[index] < target {
		if index == start {
			if nums[end] == target {
				return end
			} else {
				return -1
			}
		} else {
			return searchNums(nums,index,end,target)
		}
	} else {
		return index
	}
}

func searchBefore(nums []int,start,end,target int) int {
	if start == end {
		return -1
	}
	index := (start + end)/2
	if nums[index] > target {
		return searchBefore(nums,start,index,target)
	} else if nums[index] < target {
		if index == start {
			if nums[end] == target {
				return end
			} else {
				return -1
			}
		} else {
			return searchBefore(nums,index,end,target)
		}
	} else {
		if index > 0 && nums[index-1] == target{
			return searchBefore(nums, start, index, target)
		} else {
			return index
		}
	}
}

func searchAfter(nums []int,start,end,target int) int {
	if start == end {
		return -1
	}
	index := (start + end)/2
	if nums[index] > target {
		return searchAfter(nums,start,index,target)
	} else if nums[index] < target {
		if index == start {
			if nums[end] == target {
				return end
			} else {
				return -1
			}
		} else {
			return searchAfter(nums,index,end,target)
		}
	} else {
		if index < len(nums)-1 && nums[index+1] == target{
			if index == start {
				if nums[end] == target {
					return end
				} else {
					return -1
				}
			} else {
				return searchAfter(nums,index,end,target)
			}
		} else {
			return index
		}
	}
}