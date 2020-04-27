package medium

import (
	"sort"
	"strconv"
)

/*
给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用一次。

说明：

所有数字（包括目标数）都是正整数。
解集不能包含重复的组合。 
示例 1:

输入: candidates = [10,1,2,7,6,1,5], target = 8,
所求解集为:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
示例 2:

输入: candidates = [2,5,2,1,2], target = 5,
所求解集为:
[
  [1,2,2],
  [5]
]
*/
func combinationSum2(candidates []int, target int) [][]int {
	tempResult := make([]int,0)
	result := make([][]int,0)
	resMap := make(map[string]string)
	tranckBackForCombinationSum2(candidates,target,&tempResult,&result,0,resMap)
	return result
}

func tranckBackForCombinationSum2(candidates []int, target int,tempResult *[]int,result *[][]int,index int,resMap map[string]string) {
	for i:=index;i<len(candidates);i++ {
		if candidates[i] == target {
			*tempResult = append(*tempResult,candidates[i])
			item,key := copyArrayForCombinationSum2(*tempResult)
			if _,exist := resMap[key];!exist {
				*result = append(*result,item)
				resMap[key] = ""
			}
			*tempResult = (*tempResult)[:len(*tempResult)-1]
		} else if candidates[i] < target {
			target = target - candidates[i]
			*tempResult = append(*tempResult,candidates[i])
			tranckBackForCombinationSum2(candidates,target,tempResult,result,i+1,resMap)
			target = target + candidates[i]
			*tempResult = (*tempResult)[:len(*tempResult)-1]

		}
	}
}

func copyArrayForCombinationSum2 (array []int) ([]int,string) {
	item := make([]int,0)
	key := ""
	for i:=0;i<len(array);i++ {
		item = append(item,array[i])
	}
	sort.Ints(item)
	for i:=0;i<len(item);i++ {
		key = key + "_" + strconv.Itoa(item[i])
	}
	return item,key
}