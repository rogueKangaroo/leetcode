package medium

import (
	"fmt"
	"testing"
	"time"

	"github.com/luci/go-render/render"
)

func Test_splitWord(t *testing.T) {
	word := "ABCDEF"
	got, got1 := splitWord(word)
	fmt.Println(fmt.Sprintf("splitWord result is %v %v", render.Render(got), render.Render(got1)))
	if got == 'A' {
		fmt.Println(fmt.Sprintf("splitWord result equal"))
	}
	time.Sleep(1 * time.Second)
}

/*
board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

给定 word = "ABCCED", 返回 true
给定 word = "SEE", 返回 true
给定 word = "ABCB", 返回 false

*/
func Test_exist(t *testing.T) {
	board := [][]byte{
		{'A','B','C','E'},
		{'S','F','C','S'},
		{'A','D','E','E'},
	}
	word := "ABCCED"
	got := exist(board, word)
	fmt.Println(fmt.Sprintf("exist result is %v", render.Render(got)))
	time.Sleep(1 * time.Second)
}

func Test_exist1(t *testing.T) {
	board := [][]byte{
		{'A','B','C','E'},
		{'S','F','C','S'},
		{'A','D','E','E'},
	}
	word := "SEE"
	got := exist(board, word)
	fmt.Println(fmt.Sprintf("exist result is %v", render.Render(got)))
	time.Sleep(1 * time.Second)
}

func Test_exist2(t *testing.T) {
	board := [][]byte{
		{'A','B','C','E'},
		{'S','F','C','S'},
		{'A','D','E','E'},
	}
	word := "ABCB"
	got := exist(board, word)
	fmt.Println(fmt.Sprintf("exist result is %v", render.Render(got)))
	time.Sleep(1 * time.Second)
}
