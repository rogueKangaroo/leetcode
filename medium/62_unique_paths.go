package medium

/*
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。

问总共有多少条不同的路径？



例如，一个7 x 3 的网格。有多少可能的路径？

说明：m 和 n 的值均不超过 100。

示例 1:

	输入: m = 3, n = 2
	输出: 3
	解释:
	从左上角开始，总共有 3 条路径可以到达右下角。
	1. 向右 -> 向右 -> 向下
	2. 向右 -> 向下 -> 向右
	3. 向下 -> 向右 -> 向右

示例 2:

	输入: m = 7, n = 3
	输出: 28

解法：
	1.公式法
	res = C(m+n-2)(n-1)

	2.dp
	f(n) = num[i,j] = num[i-1,j]+num[i,j-1] // i>0.j>0
	f(n) = num[i,0] = 1
	f(n) = num[0,j] = 1
*/
func uniquePaths(m int, n int) int {
	mo := m + n - 2
	su := 0

	if m > n {
		su = n - 1
	} else {
		su = m - 1
	}

	mo_max := mo
	su_max := su
	if su_max == 0 {
		return 1
	}
	for i := 1; i < su; i++ {
		mo_max = mo_max * (mo - i)
		su_max = su_max * (su - i)
	}
	return mo_max / su_max
}

func uniquePaths2(m int, n int) int {
	m_array := make([][]int, m)
	for i := 0; i < m; i++ {
		n_array := make([]int, n)
		for j := 0; j < n; j++ {
			n_array[j] = 1
		}
		m_array[i] = n_array
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			m_array[i][j] = m_array[i-1][j] + m_array[i][j-1]
		}
	}
	return m_array[m-1][n-1]
}
