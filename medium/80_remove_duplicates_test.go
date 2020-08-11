package medium

import (
	"fmt"
	"github.com/luci/go-render/render"
	"testing"
	"time"
)

func Test_removeDuplicates(t *testing.T) {
	nums := []int{1,1,1,2,2,3}
	got := removeDuplicates(nums)
	fmt.Println(fmt.Sprintf("removeDuplicates result is %v", render.Render(got)))
	fmt.Println(fmt.Sprintf("removeDuplicates nums result is %v", render.Render(nums[:got])))
	time.Sleep(1 * time.Second)
}


func Test_removeDuplicates1(t *testing.T) {
	nums := []int{0,0,1,1,1,1,2,3,3}
	got := removeDuplicates(nums)
	fmt.Println(fmt.Sprintf("removeDuplicates result is %v", render.Render(got)))
	fmt.Println(fmt.Sprintf("removeDuplicates nums result is %v", render.Render(nums[:got])))
	time.Sleep(1 * time.Second)
}

func Test_removeDuplicates2(t *testing.T) {
	nums := []int{1,1,1}
	got := removeDuplicates(nums)
	fmt.Println(fmt.Sprintf("removeDuplicates result is %v", render.Render(got)))
	fmt.Println(fmt.Sprintf("removeDuplicates nums result is %v", render.Render(nums[:got])))
	time.Sleep(1 * time.Second)
}