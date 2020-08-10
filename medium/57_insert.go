package medium

/*
给出一个无重叠的 ，按照区间起始端点排序的区间列表。

在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。

 

示例 1：

输入：intervals = [[1,3],[6,9]], newInterval = [2,5]
输出：[[1,5],[6,9]]
示例 2：

输入：intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
输出：[[1,2],[3,10],[12,16]]
解释：这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。

*/
func insert(intervals [][]int, newInterval []int) [][]int {
	isInsert := false
	for i:=0;i<len(intervals);i++ {
		if intervals[i][0] > newInterval[0] {
			intervals = append(intervals,intervals[len(intervals)-1])
			for j:=len(intervals)-2;j>i;j--{
				intervals[j] = intervals[j-1]
			}
			intervals[i] = newInterval
			isInsert = true
			break
		}
	}
	if !isInsert {
		intervals = append(intervals,newInterval)
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

func checkMerge_insert(item []int, checkItem []int) (bool,[]int) {
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