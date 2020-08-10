package medium

import (
	"fmt"
	"testing"
	"time"

	"github.com/luci/go-render/render"
)

func Test_binarySearch(t *testing.T) {
	line := []int{10, 11, 16, 20}
	target := 13
	got := binarySearch(line, target)
	fmt.Println(fmt.Sprintf("binarySearch result is %v", render.Render(got)))
	time.Sleep(1 * time.Second)
}

/*
输入:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 3
输出: true
*/
func Test_searchMatrix(t *testing.T) {
	matrix := [][]int{
		{1,   3,  5,  7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}
	target := 3
	got := searchMatrix(matrix, target)
	fmt.Println(fmt.Sprintf("searchMatrix result is %v", render.Render(got)))
	time.Sleep(1 * time.Second)
}

/*
输入:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 13
输出: false
*/
func Test_searchMatrix1(t *testing.T) {
	matrix := [][]int{
		{1,   3,  5,  7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}
	target := 13
	got := searchMatrix(matrix, target)
	fmt.Println(fmt.Sprintf("searchMatrix result is %v", render.Render(got)))
	time.Sleep(1 * time.Second)
}

/*
输入:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 13
输出: false
*/
func Test_searchMatrix2(t *testing.T) {
	matrix := [][]int{}
	target := 1
	got := searchMatrix(matrix, target)
	fmt.Println(fmt.Sprintf("searchMatrix result is %v", render.Render(got)))
	time.Sleep(1 * time.Second)
}