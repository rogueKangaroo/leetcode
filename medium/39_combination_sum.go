package medium

/*
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：

所有数字（包括 target）都是正整数。
解集不能包含重复的组合。 
示例 1:

输入: candidates = [2,3,6,7], target = 7,
所求解集为:
[
  [7],
  [2,2,3]
]
示例 2:

输入: candidates = [2,3,5], target = 8,
所求解集为:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]

解法：回溯法：for + 递归，整除剪枝
*/
func combinationSum(candidates []int, target int) [][]int {
	tempResult := make([]int,0)
	result := make([][]int,0)
	tranckBackForCombinationSum(candidates,target,&tempResult,&result,0)
	return result
}

func tranckBackForCombinationSum(candidates []int, target int,tempResult *[]int,result *[][]int,index int) {
	for i:=index;i<len(candidates);i++ {
		if candidates[i] == target {
			*tempResult = append(*tempResult,candidates[i])
			item := copyArray(*tempResult)
			*result = append(*result,item)
			*tempResult = (*tempResult)[:len(*tempResult)-1]
		} else if candidates[i] < target {
			target = target - candidates[i]
			*tempResult = append(*tempResult,candidates[i])
			tranckBackForCombinationSum(candidates,target,tempResult,result,i)
			target = target + candidates[i]
			*tempResult = (*tempResult)[:len(*tempResult)-1]

		}
	}
}

func copyArray (array []int) []int {
	item := make([]int,0)
	for i:=0;i<len(array);i++ {
		item = append(item,array[i])
	}
	return item
}