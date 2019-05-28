package medium

/*
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。

现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？

网格中的障碍物和空位置分别用 1 和 0 来表示。

说明：m 和 n 的值均不超过 100。

示例 1:

输入:
	[
	  [0,0,0],
	  [0,1,0],
	  [0,0,0]
	]
输出: 2
解释:
	3x3 网格的正中间有一个障碍物。
	从左上角到右下角一共有 2 条不同的路径：
	1. 向右 -> 向右 -> 向下 -> 向下
	2. 向下 -> 向下 -> 向右 -> 向右

解法：
	dp:
	f(n) = num[i,j] = num[i-1,j] + num[i,j-1] // input[i][j] != 1
	f(n) = num[i,j] = 0                       // input[i][j] == 1
*/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	result := make([][]int, len(obstacleGrid))
	for i := 0; i < len(obstacleGrid); i++ {
		row := make([]int, len(obstacleGrid[0]))
		copy(row, obstacleGrid[i])
		result[i] = row
	}
	i_flag := true
	j_flag := true
	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < len(obstacleGrid[0]); j++ {
			if i == 0 {
				if obstacleGrid[i][j] == 0 && j_flag {
					result[i][j] = 1
				} else {
					j_flag = false
					result[i][j] = 0
				}
			}
			if j == 0 {
				if obstacleGrid[i][j] == 0 && i_flag {
					result[i][j] = 1
				} else {
					i_flag = false
					result[i][j] = 0
				}
			}
			if i > 0 && j > 0 {
				if obstacleGrid[i][j] == 0 {
					result[i][j] = result[i-1][j] + result[i][j-1]
				} else {
					result[i][j] = 0
				}
			}
		}
	}
	return result[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}
