package medium

import (
	"fmt"
	"github.com/luci/go-render/render"
	"testing"
	"time"
)

/*
输入: [2,0,2,1,1,0]
输出: [0,0,1,1,2,2]
*/
func Test_sortColors(t *testing.T) {
	nums := []int{2,0,2,1,1,0}
	sortColors(nums)
	fmt.Println(fmt.Sprintf("sortColors result is %v", render.Render(nums)))
	time.Sleep(1 * time.Second)
}

func Test_sortColors1(t *testing.T) {
	nums := []int{2,0,2,1,1,0,2,1,0,2,1,2,1,2,2,0,0}
	sortColors(nums)
	fmt.Println(fmt.Sprintf("sortColors result is %v", render.Render(nums)))
	time.Sleep(1 * time.Second)
}

func Test_sortColors2(t *testing.T) {
	nums := []int{2,0,1}
	sortColors(nums)
	fmt.Println(fmt.Sprintf("sortColors result is %v", render.Render(nums)))
	time.Sleep(1 * time.Second)
}

func Test_sortColors3(t *testing.T) {
	nums := []int{2}
	sortColors(nums)
	fmt.Println(fmt.Sprintf("sortColors result is %v", render.Render(nums)))
	time.Sleep(1 * time.Second)
}

func Test_sortColors4(t *testing.T) {
	nums := []int{0,1,0}
	sortColors(nums)
	fmt.Println(fmt.Sprintf("sortColors result is %v", render.Render(nums)))
	time.Sleep(1 * time.Second)
}