package medium

import (
	"fmt"
	"testing"
	"time"
)

func Test_canJump(t *testing.T) {
	nums := []int{2,3,1,1,4}
	got := canJump(nums)
	fmt.Println(fmt.Sprintf("canJump result is %v", got))
	time.Sleep(1 * time.Second)

	nums = []int{3,2,1,0,4}
	got = canJump(nums)
	fmt.Println(fmt.Sprintf("canJump result is %v", got))
	time.Sleep(1 * time.Second)

	nums = []int{2,0,0}
	got = canJump(nums)
	fmt.Println(fmt.Sprintf("canJump result is %v", got))
	time.Sleep(1 * time.Second)
}
