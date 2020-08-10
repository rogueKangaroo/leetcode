package medium

/*
给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

输入: nums = [1,2,3]
输出:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
解法：简单递归
*/
func subsets(nums []int) [][]int {
	result := make([][]int,0)
	result = append(result,[]int{})
	if len(nums) == 0 {
		return result
	}
	for i:=1;i<=len(nums);i++{
		item := make([]int,i)
		recursionCombineSubsets(len(nums),i,len(nums),i,0,&result,0,0,item,nums)
	}
	return result
}

func recursionCombineSubsets(n,k,nFlag,kFlag,start int,result *[][]int,index,length int,item []int,nums []int)  {
	if length == kFlag {
		*result = append(*result,copyItemSubsets(item))
		return
	}
	if k <= 0 || n < k {
		return
	}
	for i:=start;i<len(nums);i++ {
		item[index] = nums[i]
		length ++
		recursionCombineSubsets(n-1,k-1,nFlag,kFlag,i+1,result,index + 1,length,item,nums)
		length --
	}
}

func copyItemSubsets(item []int) []int {
	newItem := make([]int,0)
	for i:=0;i<len(item);i++{
		newItem = append(newItem,item[i])
	}
	return newItem
}