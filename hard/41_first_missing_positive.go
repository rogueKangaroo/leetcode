package hard

/*
给你一个未排序的整数数组，请你找出其中没有出现的最小的正整数。

示例 1:

输入: [1,2,0]
输出: 3
示例 2:

输入: [3,4,-1,1]
输出: 2
示例 3:

输入: [7,8,9,11,12]
输出: 1
 

提示：

你的算法的时间复杂度应为O(n)，并且只能使用常数级别的额外空间。
解法：先将所有负数和0变成1，利用下标标识某个数字是否存在，如果存在则改为负值

*/
func firstMissingPositive(nums []int) int {
	containsOne := 0

	for i:=0;i<len(nums);i++ {
		if nums[i] == 1 {
			containsOne ++
		} else if nums[i] <=0 {
			nums[i] = 1
		}
	}
	if containsOne == 0 {
		return 1
	}
	for i:=0;i<len(nums);i++{
		if absNums(nums[i]) <= len(nums) {
			if nums[i] >= 0 && nums[nums[i]-1] > 0{
				nums[nums[i]-1] = 0 - nums[nums[i]-1]
			} else if nums[i] < 0 && nums[0-(nums[i]+1)] > 0{
				nums[0-(nums[i]+1)]	= 0 - nums[0-(nums[i]+1)]
			}

		}
	}

	for i:=0;i<len(nums);i++{
		if nums[i] > 0 {
			return  i + 1
		}
	}
	return len(nums) + 1
}

func absNums(num int) int {
	if num < 0 {
		return 0 - num
	}else {
		return num
	}
}