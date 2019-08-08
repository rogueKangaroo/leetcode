package medium

/**
给你一个整数数组 nums，每次 操作 会从中选择一个元素并 将该元素的值减少 1。

如果符合下列情况之一，则数组 A 就是 锯齿数组：

每个偶数索引对应的元素都大于相邻的元素，即 A[0] > A[1] < A[2] > A[3] < A[4] > ...
或者，每个奇数索引对应的元素都大于相邻的元素，即 A[0] < A[1] > A[2] < A[3] > A[4] < ...
返回将数组 nums 转换为锯齿数组所需的最小操作次数。

 

示例 1：

输入：nums = [1,2,3]
输出：2
解释：我们可以把 2 递减到 0，或把 3 递减到 1。
示例 2：

输入：nums = [9,6,1,6,2]
输出：4
 
1，2，3，4，5，6
提示：

1 <= nums.length <= 1000
1 <= nums[i] <= 1000

10,4,4,10,10,6,2,3
**/
func movesToMakeZigzag(nums []int) int {
	sum_j := 0
	sum_o := 0
	last_j := nums[0]
	last_o := nums[0]
	for i:=1;i<len(nums);i++ {
		isSum_o := i%2 == 0
		// 如果是偶数位
		if isSum_o {
			if nums[i] <= last_o {
				sum_o = sum_o + (last_o - nums[i]) + 1
				last_o = nums[i]
			}
			last_o = nums[i]
			if nums[i] >= last_j {
				sum_j = sum_j + (nums[i] - last_j) + 1
				last_j = last_j - 1
			} else {
				last_j = nums[i]
			}
		} else {
			if nums[i] >= last_o {
				sum_o = sum_o + (nums[i] - last_o) + 1
				last_o = last_o - 1
			} else {
				last_o = nums[i]
			}
			if nums[i] <= last_j {
				sum_j = sum_j + (last_j - nums[i]) + 1
			}
			last_j = nums[i]
		}
	}
	if sum_o > sum_j {
		return sum_j
	} else {
		return sum_o
	}
}