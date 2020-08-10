package medium

/*
编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：

每行中的整数从左到右按升序排列。
每行的第一个整数大于前一行的最后一个整数。
示例 1:

输入:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 3
输出: true
示例 2:

输入:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 13
输出: false

解法：确认行标，二分查找
*/
func searchMatrix(matrix [][]int, target int) bool {
	for i:=0;i<len(matrix);i++ {
		if len(matrix[i]) > 0 && matrix[i][0] <= target && matrix[i][len(matrix[i])-1] >= target {
			return binarySearch(matrix[i],target)
		}
	}
	return false
}

func binarySearch(line []int,target int) bool {
	start := 0
	end := len(line) - 1
	for {
		if target > line[start] {
			index := (start + end) / 2
			if index == start {
				index = end
				start = index
			}
			if line[index] > target {
				end = index
			} else if line[index] < target {
				start = index
			} else {
				return true
			}
			if start == end {
				return false
			}
		} else if target == line[start]{
			return true
		} else {
			return false
		}
	}
}
