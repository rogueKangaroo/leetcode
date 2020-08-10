package medium

import (
	"fmt"
	"github.com/luci/go-render/render"
	"testing"
	"time"
)

/*
示例 1：

输入：intervals = [[1,3],[6,9]], newInterval = [2,5]
输出：[[1,5],[6,9]]
示例 2：

输入：intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
输出：[[1,2],[3,10],[12,16]]
*/
func Test_insert(t *testing.T) {
	intervals := [][]int{{1,3},{6,9}}
	newInterval := []int{2,5}
	got := insert(intervals, newInterval)
	fmt.Println(fmt.Sprintf("insert result is %v", render.Render(got)))
	time.Sleep(1 * time.Second)
}

func Test_insert1(t *testing.T) {
	intervals := [][]int{{1,2},{3,5},{6,7},{8,10},{12,16}}
	newInterval := []int{4,8}
	got := insert(intervals, newInterval)
	fmt.Println(fmt.Sprintf("insert result is %v", render.Render(got)))
	time.Sleep(1 * time.Second)
}