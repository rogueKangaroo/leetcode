package medium

/*
给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

示例:

输入:
	[
	  [1,3,1],
	  [1,5,1],
	  [4,2,1]
	]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。

解法：
	dp:
	f(n) = number[i,j] = input[0,0]  // i==0 && j==0
	f(n) = number[i,j] = number[i-1,0] + input[i,j] // i>0 && j == 0
	f(n) = number[i,j] = number[0,j-1] + input[i,j] // j>0 && i == 0
    f(n) = number[i,j] = min(number[i,j-1],number[i-1,j]) + input[i,j] // i>0 && j>0
*/

func minPathSum(grid [][]int) int {
	number := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		row := make([]int, len(grid[0]))
		copy(row, grid[i])
		number[i] = row
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j == 0 {
				number[i][j] = grid[i][j]
			}
			if i > 0 && j == 0 {
				number[i][j] = number[i-1][j] + grid[i][j]
			}
			if j > 0 && i == 0 {
				number[i][j] = number[i][j-1] + grid[i][j]
			}
			if i > 0 && j > 0 {
				number[i][j] = getMin(number[i][j-1], number[i-1][j]) + grid[i][j]
			}
		}
	}
	return number[len(grid)-1][len(grid[0])-1]
}

func getMin(m, n int) int {
	if m > n {
		return n
	}
	return m
}
