package hard

/*

编写一个程序，通过已填充的空格来解决数独问题。

一个数独的解法需遵循如下规则：

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
空白格用 '.' 表示。



一个数独。



答案被标成红色。

Note:

给定的数独序列只包含数字 1-9 和字符 '.' 。
你可以假设给定的数独只有唯一解。
给定数独永远是 9x9 形式的。

解法：回溯法
*/
func solveSudoku(board [][]byte)  {
	backTrackForSolveSudoku(&board,0,0)
}

func backTrackForSolveSudoku(board *[][]byte,row,col int)  {
	newRow,newCol := getRowAndCol(board,row,col)
	if  newRow == -1 && newCol == -1 {
		return
	}

	for i:=1;i<=9;i++ {
		if isValidForSudoku(newCol,newRow,byte(i),board) {
			(*board)[newRow][newCol] = byte(i)
			backTrackForSolveSudoku(board,row,col)
			(*board)[newRow][newCol] = '.'
		}
	}
}

func isValidForSudoku(row,col int, num byte,board *[][]byte) bool {
	for i:=0;i<9;i++ {
		if (*board)[i][col] == num || (*board)[row][i] == num {
			return false
		}
	}
	tempRow := row/3 * 3
	tempCol := row/3 * 3
	for i:=0;i<3;i++{
		for j:=0;j<3;j++ {
			if (*board)[tempCol+i][tempRow+j] == num {
				return false
			}
		}
	}
	return true
}
func getRowAndCol(board *[][]byte,row,col int) (int,int) {
	for i:=row;i<9;i++{
		for j:=col;j<9;j++{
			if (*board)[i][j] == '.' {
				return i,j
			}
		}
	}
	return -1,-1
}
