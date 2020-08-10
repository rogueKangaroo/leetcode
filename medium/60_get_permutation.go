package medium

import "strconv"

/*
给出集合 [1,2,3,…,n]，其所有元素共有 n! 种排列。

按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：

"123"
"132"
"213"
"231"
"312"
"321"
给定 n 和 k，返回第 k 个排列。

说明：

给定 n 的范围是 [1, 9]。
给定 k 的范围是[1,  n!]。
示例 1:

输入: n = 3, k = 3
输出: "213"
示例 2:

输入: n = 4, k = 9
输出: "2314"
"1234"
"1243"
"1324"
"1342"
"1423"
"1432"
"2134"
"2143"
"2314"
"2341"
"2413"
"2431"
"3124"
"3142"
"3214"
"3241"
"3412"
"3421"
"4123"
"4132"
"4213"
"4231"
"4312"
"4321"
*/
func getPermutation(n int, k int) string {
	result := ""
	permutationArray := make([]int,n)
	for i:=1;i<=n;i++{
		permutationArray[i-1] = i
	}
	getPermutationSort(permutationArray,n,k,&result)
	return result
}

func getPermutationSort(nArray []int, n,k int, result *string)  {
	if len(*result) == len(nArray){
		return
	} else if len(nArray) - len(*result) == 1{
		for i:=0;i<len(nArray);i++{
			if nArray[i] != 0 {
				*result = *result + strconv.Itoa(nArray[i])
				nArray[i] = 0
				break
			}
		}
		return
	} else {
		flag := 1
		for i:=1;i<n;i++{
			flag = flag * i
		}
		num := k/flag
		if k%flag == 0 {
			num --
		}
		flagNum := num
		for i:=0;i<len(nArray);i++{
			if nArray[i] != 0 {
				if num > 0 {
					num --
				} else {
					*result = *result + strconv.Itoa(nArray[i])
					nArray[i] = 0
					break
				}
			}
		}
		getPermutationSort(nArray,n-1,k-flag*flagNum,result)
	}
}
