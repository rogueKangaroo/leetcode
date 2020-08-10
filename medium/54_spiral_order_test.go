package medium

import (
	"fmt"
	"github.com/luci/go-render/render"
	"testing"
	"time"
)

/*
示例 1:

输入:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
输出: [1,2,3,6,9,8,7,4,5]
示例 2:

输入:
[
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9,10,11,12]
]
输出: [1,2,3,4,8,12,11,10,9,5,6,7]
*/
func Test_spiralOrder(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	got := spiralOrder(matrix)
	fmt.Println(fmt.Sprintf("spiralOrder result is %v", render.Render(got)))
	time.Sleep(1 * time.Second)
}

func Test_spiralOrder1(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9,10,11,12},
	}
	got := spiralOrder(matrix)
	fmt.Println(fmt.Sprintf("spiralOrder result is %v", render.Render(got)))
	time.Sleep(1 * time.Second)
}