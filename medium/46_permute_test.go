package medium

import (
	"fmt"
	"testing"
	"time"
)

// 输入: [1,2,3]
//输出:
//[
//  [1,2,3],
//  [1,3,2],
//  [2,1,3],
//  [2,3,1],
//  [3,1,2],
//  [3,2,1]
//]
func Test_permute(t *testing.T) {
	nums := []int{1,2,3}
	got := permute(nums)
	fmt.Println(fmt.Sprintf("permute result is %v", got))
	time.Sleep(1 * time.Second)
}

func Test_permute1(t *testing.T) {
	nums := []int{1,2,3,4}
	got := permute(nums)
	fmt.Println(fmt.Sprintf("permute result is %v", got))
	time.Sleep(1 * time.Second)
}