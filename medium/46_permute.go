package medium

/*
给定一个 没有重复 数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
解法：回溯法 for + 递归
*/
func permute(nums []int) [][]int {
	itemRes := []int{}
	numsMap := make(map[int]string)
	res := [][]int{}
	tranckBackForPermute(nums,numsMap,&itemRes,&res)
	return res
}

func tranckBackForPermute(nums []int,numsMap map[int]string,itemRes *[]int,res *[][]int)  {
	if len(*itemRes) == len(nums) {
		item:= copyIntArray(*itemRes)
		*res = append(*res,item)
		return
	}

	for i:=0;i<len(nums);i++ {
		if value,exist := numsMap[nums[i]];!exist || (exist&&value == ""){
			numsMap[nums[i]] = "1"
			*itemRes = append(*itemRes,nums[i])
			tranckBackForPermute(nums,numsMap,itemRes,res)
			numsMap[nums[i]] = ""
			*itemRes = (*itemRes)[:len(*itemRes)-1]
		}
	}
}

func copyIntArray(nums []int) []int {
	res := make([]int,0)
	for i:=0;i<len(nums);i++{
		res = append(res, nums[i])
	}
	return res
}