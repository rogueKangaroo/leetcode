package medium

import (
	"sort"
	"strconv"
)

/*
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

注意：

答案中不可以包含重复的四元组。

示例：

给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
解法：先排序，外层遍历+内层双指针
*/
func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return [][]int{}
	}
	sort.Ints(nums)
	resultArrays := make([][]int,0)
	reslutMap := make(map[string]string)
	for i:=0;i<len(nums)-3;i++{
		for j:= i + 1;j<len(nums) - 2;j++ {
			start := j + 1
			end := len(nums) - 1
			for {
				if start == end {
					break
				}
				sum := nums[i] + nums[j] + nums[start] + nums[end]
				if sum < target {
					start ++
				} else if sum > target {
					end --
				} else {
					tempStr := strconv.Itoa(nums[i]) + strconv.Itoa(nums[j]) + strconv.Itoa(nums[start]) + strconv.Itoa(nums[end])
					if _,exist := reslutMap[tempStr];!exist {
						temp := []int{nums[i],nums[j],nums[start],nums[end]}
						resultArrays = append(resultArrays,temp)
						reslutMap[tempStr] = ""
					}
					start ++
				}
			}
		}
	}
	return resultArrays
}
