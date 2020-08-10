package medium

/*
给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

示例:

输入: n = 4, k = 2
输出:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
解法：递归解法：F(n,k) = [c[0],F(n-1,k-1)] & F(n-1,k)
*/
func combine(n int, k int) [][]int {
	result := make([][]int,0)
	item := make([]int,k)
	if k == 1 {
		for i:=1;i<=n;i++{
			newItem := []int{i}
			result = append(result,newItem)
		}
	} else {
		recursionCombine(n,k,n,k,1,&result,0,0,item)
	}
	return result
}

func recursionCombine(n,k,nFlag,kFlag,start int,result *[][]int,index,length int,item []int)  {
	if length == kFlag {
		*result = append(*result,copyItem(item))
		return
	}
	if k <= 0 || n < k || nFlag - start + 1 < kFlag - index{
		return
	}
	for i:=start;i<=nFlag;i++ {
		item[index] = i
		length ++
		recursionCombine(n-1,k-1,nFlag,kFlag,i+1,result,index + 1,length,item)
		length --
	}
}

func copyItem(item []int) []int {
	newItem := make([]int,0)
	for i:=0;i<len(item);i++{
		newItem = append(newItem,item[i])
	}
	return newItem
}