package medium

/*
给出一个区间的集合，请合并所有重叠的区间。

示例 1:

输入: [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2:

输入: [[1,4],[4,5]]
输出: [[1,5]]
解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
*/
func merge(intervals [][]int) [][]int {
	for i:=0;i<len(intervals);i++{
		for j:=len(intervals)-1;j>i;j--{
			if intervals[j][0] < intervals[j-1][0] {
				temp := intervals[j]
				intervals[j] = intervals[j-1]
				intervals[j-1] = temp
			}
		}
	}
	resultArray := make([][]int,0)
	index := 0
	for i:=0;i<len(intervals);i++{
		if i == 0 {
			resultArray = append(resultArray,intervals[i])
			index ++
		} else {
			isNeedMerge,resultItem := checkMerge(resultArray[index-1],intervals[i])
			if isNeedMerge {
				resultArray[index-1] = resultItem
			} else {
				resultArray = append(resultArray,intervals[i])
				index ++
			}
		}
	}
	return resultArray
}

func checkMerge(item []int, checkItem []int) (bool,[]int) {
	if item[1] >= checkItem[0] {
		left := item[0]
		if left > checkItem[0]{
			left = checkItem[0]
		}
		right := checkItem[1]
		if right < item[1]{
			right = item[1]
		}
		return true,[]int{left,right}
	} else {
		return false,[]int{}
	}
}

