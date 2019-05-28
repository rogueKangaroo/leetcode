package hard

/*
给定一个仅包含 0 和 1 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。

示例:

输入:
[
  ["1","0","1","0","0"],
  ["1","0","1","1","1"],
  ["1","1","1","1","1"],
  ["1","0","0","1","0"]
]
输出: 6

解法：
	转换成柱状图最大矩形面积的问题来求解
	柱状图最大矩形面积采用单调栈来求解
*/
func maximalRectangle(matrix [][]byte) int {
	maxRectangle := 0
	for i := 0; i < len(matrix); i++ {
		x := make([]int, len(matrix[i]))
		for k := 0; k < len(matrix[i]); k++ {
			for j := i; j >= 0; j-- {
				if matrix[j][k] == '1' {
					x[k] = x[k] + 1
				} else {
					break
				}
			}
		}
		max := getMaxRectangle(x)
		if max > maxRectangle {
			maxRectangle = max
		}
	}
	return maxRectangle
}

func getMaxRectangle(x []int) int {
	stack := make([]int, len(x))
	index := 0
	maxRectangle := 0
	for i := 0; i < len(x); i++ {
		if index == 0 || x[stack[index-1]] < x[i] {
			stack[index] = i
			index++
		} else {
			index--
			width := 0
			if index == 0 {
				width = i
			} else {
				width = i - stack[index-1] - 1
			}
			rectangle := width * x[stack[index]]
			if rectangle > maxRectangle {
				maxRectangle = rectangle
			}

			i--
		}
	}
	for i := index - 1; i >= 0; i-- {
		index--
		width := 0
		if index == 0 {
			width = len(x)
		} else {
			width = len(x) - stack[index-1] - 1
		}
		rectangle := width * x[stack[index]]
		if rectangle > maxRectangle {
			maxRectangle = rectangle
		}
	}
	return maxRectangle
}
