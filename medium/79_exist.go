package medium

/*
给定一个二维网格和一个单词，找出该单词是否存在于网格中。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

 

示例:

board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

给定 word = "ABCCED", 返回 true
给定 word = "SEE", 返回 true
给定 word = "ABCB", 返回 false
 

提示：

board 和 word 中只包含大写和小写英文字母。
1 <= board.length <= 200
1 <= board[i].length <= 200
1 <= word.length <= 10^3

解法：简单递归配对
*/
func exist(board [][]byte, word string) bool {
	firstByte,wordStr := splitWord(word)
	for i:=0;i<len(board);i++{
		for j:=0;j<len(board[0]);j++{
			if board[i][j] == firstByte{
				board[i][j] = ' '
				if pathExist(board,wordStr,i,j) {
					return true
				} else {
					board[i][j] = firstByte
				}
			}
		}
	}
	return false
}

func pathExist(board [][]byte, word string,i,j int) bool {
	if word == ""{
		return true
	}
	firstByte,wordStr := splitWord(word)
	leftI,leftJ := i,j-1
	if leftJ >= 0 && board[leftI][leftJ] == firstByte{
		board[leftI][leftJ] = ' '
		if pathExist(board,wordStr,leftI,leftJ){
			return true
		} else {
			board[leftI][leftJ] = firstByte
		}
	}

	rightI,rightJ := i,j+1
	if rightJ <= len(board[0]) - 1 && board[rightI][rightJ] == firstByte{
		board[rightI][rightJ] = ' '
		if pathExist(board,wordStr,rightI,rightJ){
			return true
		} else {
			board[rightI][rightJ] = firstByte
		}
	}

	upI,upJ := i-1,j
	if upI >=0 && board[upI][upJ] == firstByte{
		board[upI][upJ] = ' '
		if pathExist(board,wordStr,upI,upJ){
			return true
		} else {
			board[upI][upJ] = firstByte
		}
	}

	downI,downJ := i+1,j
	if downI <= len(board) - 1 && board[downI][downJ] == firstByte{
		board[downI][downJ] = ' '
		if pathExist(board,wordStr,downI,downJ){
			return true
		} else {
			board[downI][downJ] = firstByte
		}
	}

	return false
}

func splitWord(word string) (byte,string) {
	wordByte := []byte(word)
	if len(wordByte) == 1 {
		return wordByte[0],""
	}
	firstByte := wordByte[0]
	wordStr := string(wordByte[1:])
	return firstByte,wordStr
}
