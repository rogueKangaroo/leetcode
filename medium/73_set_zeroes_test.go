package medium

import (
	"fmt"
	"github.com/luci/go-render/render"
	"testing"
	"time"
)

/*
输入:
[
  [1,1,1],
  [1,0,1],
  [1,1,1]
]
输出:
[
  [1,0,1],
  [0,0,0],
  [1,0,1]
]
*/
func Test_setZeroes(t *testing.T) {
	matrix := [][]int{
		{1,1,1},
		{1,0,1},
		{1,1,1},
	}
	setZeroes(matrix)
	fmt.Println(fmt.Sprintf("setZeroes result is"))
	for i:=0;i<len(matrix);i++{
		fmt.Println(fmt.Sprintf("%v", render.Render(matrix[i])))
	}
	time.Sleep(1 * time.Second)
}

/*
输入:
[
  [0,1,2,0],
  [3,4,5,2],
  [1,3,1,5]
]
输出:
[
  [0,0,0,0],
  [0,4,5,0],
  [0,3,1,0]
]
*/
func Test_setZeroes1(t *testing.T) {
	matrix := [][]int{
		{0,1,2,0},
		{3,4,5,2},
		{1,3,1,5},
	}
	setZeroes(matrix)
	fmt.Println(fmt.Sprintf("setZeroes result is"))
	for i:=0;i<len(matrix);i++{
		fmt.Println(fmt.Sprintf("%v", render.Render(matrix[i])))
	}
	time.Sleep(1 * time.Second)
}


func Test_setZeroes2(t *testing.T) {
	matrix := [][]int{
		{0,0,2,0},
		{3,0,5,2},
		{1,3,1,5},
	}
	setZeroes(matrix)
	fmt.Println(fmt.Sprintf("setZeroes result is"))
	for i:=0;i<len(matrix);i++{
		fmt.Println(fmt.Sprintf("%v", render.Render(matrix[i])))
	}
	time.Sleep(1 * time.Second)
}