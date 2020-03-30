package medium

import (
	"fmt"
	"testing"
	"time"
)

func Test_fourSum(t *testing.T) {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	got := fourSum(nums, target)
	for i:=0;i<len(got);i++{
		fmt.Println(got[i])
	}
	time.Sleep(1*time.Second)
}

//[-3,-2,-1,0,0,1,2,3]
//0
func Test_fourSum1(t *testing.T) {
	nums := []int{-3,-2,-1,0,0,1,2,3}
	target := 0
	got := fourSum(nums, target)
	for i:=0;i<len(got);i++{
		fmt.Println(got[i])
	}
	time.Sleep(1*time.Second)
}