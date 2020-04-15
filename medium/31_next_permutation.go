package medium

import "sort"

/*
实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。

如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。

必须原地修改，只允许使用额外常数空间。

以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1
1,1,5,4,6,5 -> 1,1,5,5,4,6
1,3,5,4,3,2 -> 1,4,2,3,3,5

解法：反向遍历找到第一个拐点（由高到低），交换拐点（找到略大于拐点的数字），后面从小到大排序
*/
func nextPermutation(nums []int)  {
	flagNum := nums[len(nums)-1]
	flagIndex := len(nums) -1
	i:=len(nums)-1
	for ;i>=0;i-- {
		if nums[i] >= flagNum {
			flagNum = nums[i]
		} else {
			flagNum = nums[i]
			flagIndex = i
			break
		}
	}
	if i < 0 {
		length := len(nums)
		for i := 0; i < length/2; i++ {
			temp := nums[length-1-i]
			nums[length-1-i] = nums[i]
			nums[i] = temp
		}
		return
	}
	tempArray := nums[i+1:]
	sort.Ints(tempArray)
	j := 0
	flag := true
	for ;j<len(tempArray);j++{
		if tempArray[j] > flagNum && flag{
			temp := tempArray[j]
			tempArray[j] = flagNum
			nums[flagIndex] = temp
			flag = false
			j--
		} else {
			i ++
			nums[i] = tempArray[j]
		}
	}

}