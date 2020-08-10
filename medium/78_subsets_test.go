package medium

import (
	"fmt"
	"github.com/luci/go-render/render"
	"testing"
	"time"
)

/*
输入: nums = [1,2,3]
输出:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
*/
func Test_subsets(t *testing.T) {
	nums := []int{1,2,3}
	got := subsets(nums)
	fmt.Println(fmt.Sprintf("subsets result is %v", render.Render(got)))
	time.Sleep(1 * time.Second)
}


func Test_subsets1(t *testing.T) {
	nums := []int{2,3,4}
	got := subsets(nums)
	fmt.Println(fmt.Sprintf("subsets result is %v", render.Render(got)))
	time.Sleep(1 * time.Second)
}

func Test_subsets2(t *testing.T) {
	nums := []int{2,3,5,4}
	got := subsets(nums)
	fmt.Println(fmt.Sprintf("subsets result is %v", render.Render(got)))
	time.Sleep(1 * time.Second)
}