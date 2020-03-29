package medium

import (
	"sort"
)

/*
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.

与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).

解决方案：排序+双指针
*/
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	if len(nums) == 3{
		return nums[0] + nums[1] + nums[2]
	}

	min := nums[0] + nums[1] + nums[2]
	for i:=0;i<len(nums)-2;i++ {
		start := i+1
		end := len(nums)-1
		sum := 1<<63 - 1
		for {
			if start == end {
				break
			}
			sum = nums[start] + nums[i] + nums[end]
			if sum < target {
				start ++
			} else if sum > target {
				end --
			} else {
				return target
			}
			if MathAbs(sum - target) <  MathAbs(min - target){
				min = sum
			}
		}
	}
	return min
}

func MathAbs(num int) int {
	if num < 0 {
		return 0 - num
	} else {
		return num
	}
}