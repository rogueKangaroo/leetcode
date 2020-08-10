package medium

import (
	"fmt"
	"github.com/luci/go-render/render"
	"testing"
	"time"
)

/*
示例 1:

输入: [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2:

输入: [[1,4],[4,5]]
输出: [[1,5]]
解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
*/
func Test_merge(t *testing.T) {
	intervals := [][]int{{1,3},{2,6},{8,10},{15,18}}
	got := merge(intervals)
	fmt.Println(fmt.Sprintf("merge result is %v", render.Render(got)))
	time.Sleep(1 * time.Second)
}

func Test_merge1(t *testing.T) {
	intervals := [][]int{{1,4},{4,5}}
	got := merge(intervals)
	fmt.Println(fmt.Sprintf("merge result is %v", render.Render(got)))
	time.Sleep(1 * time.Second)
}

//[[1,4],[0,2],[3,5]]
func Test_merge2(t *testing.T) {
	intervals := [][]int{{1,4},{0,2},{3,5}}
	got := merge(intervals)
	fmt.Println(fmt.Sprintf("merge result is %v", render.Render(got)))
	time.Sleep(1 * time.Second)
}

// [[2,3],[4,5],[6,7],[8,9],[1,10]]
func Test_merge3(t *testing.T) {
	intervals := [][]int{{2,3},{4,5},{6,7},{8,9},{1,10}}
	got := merge(intervals)
	fmt.Println(fmt.Sprintf("merge result is %v", render.Render(got)))
	time.Sleep(1 * time.Second)
}