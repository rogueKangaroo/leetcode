package medium

import (
	"fmt"
	"testing"
	"time"
)

// 输入: nums = [5,7,7,8,8,10], target = 8
//输出: [3,4]
//示例 2:
//
//输入: nums = [5,7,7,8,8,10], target = 6
//输出: [-1,-1]
func Test_searchRange(t *testing.T) {
	nums := []int{5,7,7,8,8,10}
	got := searchRange(nums, 8)
	fmt.Println(fmt.Sprintf("searchRange result is %v", got))
	time.Sleep(1 * time.Second)
}


func Test_searchRange1(t *testing.T) {
	nums := []int{5,7,7,8,8,10}
	got := searchRange(nums, 6)
	fmt.Println(fmt.Sprintf("searchRange result is %v", got))
	time.Sleep(1 * time.Second)
}

func Test_searchRange2(t *testing.T) {
	nums := []int{}
	got := searchRange(nums, 6)
	fmt.Println(fmt.Sprintf("searchRange result is %v", got))
	time.Sleep(1 * time.Second)
}

func Test_searchRange3(t *testing.T) {
	nums := []int{1}
	got := searchRange(nums, 1)
	fmt.Println(fmt.Sprintf("searchRange result is %v", got))
	time.Sleep(1 * time.Second)
}

func Test_searchRange4(t *testing.T) {
	nums := []int{2,2}
	got := searchRange(nums, 2)
	fmt.Println(fmt.Sprintf("searchRange result is %v", got))
	time.Sleep(1 * time.Second)
}

//[0,0,1,1,1,2,2,3,3,3,4,4,4,4,5,5,6,6,6,8,10,10]
//4
//[10,13]
func Test_searchRange5(t *testing.T) {
	nums := []int{0,0,1,1,1,2,2,3,3,3,4,4,4,4,5,5,6,6,6,8,10,10}
	got := searchRange(nums, 4)
	fmt.Println(fmt.Sprintf("searchRange result is %v", got))
	time.Sleep(1 * time.Second)
}