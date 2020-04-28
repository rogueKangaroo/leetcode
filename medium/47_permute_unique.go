package medium

import "strconv"

/*
给定一个可包含重复数字的序列，返回所有不重复的全排列。

示例:

输入: [1,1,2]
输出:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]
*/
func permuteUnique(nums []int) [][]int {
	itemRes := []int{}
	numsMap := make(map[int]string)
	res := [][]int{}
	resMap := make(map[string]string)
	tranckBackForPermuteUnique(nums,numsMap,&itemRes,&res,resMap)
	return res
}

func tranckBackForPermuteUnique(nums []int,numsMap map[int]string,itemRes *[]int,res *[][]int, resMap map[string]string)  {
	if len(*itemRes) == len(nums) {
		item,key := copyIntArrayForPermuteUnique(*itemRes)
		if _,exist := resMap[key];!exist {
			*res = append(*res,item)
			resMap[key] = ""
		}
		return
	}

	for i:=0;i<len(nums);i++ {
		if value,exist := numsMap[i];!exist || (exist&&value == ""){
			numsMap[i] = "1"
			*itemRes = append(*itemRes,nums[i])
			tranckBackForPermuteUnique(nums,numsMap,itemRes,res,resMap)
			numsMap[i] = ""
			*itemRes = (*itemRes)[:len(*itemRes)-1]
		}
	}
}

func copyIntArrayForPermuteUnique(nums []int) ([]int,string) {
	res := make([]int,0)
	key := ""
	for i:=0;i<len(nums);i++{
		res = append(res, nums[i])
		key = key + "_" + strconv.Itoa(nums[i])
	}
	return res,key
}