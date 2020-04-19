package hard

import (
	"fmt"
	"testing"
	"time"
)

func Test_solveSudoku(t *testing.T) {
	board := [][]byte{
		{'5','3','.', '.','7','.', '.','.','.'},
		{'.','.','.', '.','.','.', '.','.','.'},
		{'.','.','.', '.','.','.', '.','.','.'},
		{},
		{},
		{},
		{},
		{},
		{},
	}
	solveSudoku(board)
	for i:=0;i<9;i++{
		fmt.Println(fmt.Sprintf("%v", board[i]))
	}
}


func Test_solveSudoku_byte(t *testing.T) {
	nums := 1
	var byteTest byte
	byteTest = byte(nums)
	fmt.Println(fmt.Sprintf("byte result is %v", byteTest))
	time.Sleep(1 * time.Second)
}