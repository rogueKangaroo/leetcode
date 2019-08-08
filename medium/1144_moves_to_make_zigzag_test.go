package medium

import (
	"fmt"
	"testing"
	"time"
)

func Test_movesToMakeZigzag(t *testing.T) {
	nums := []int{1,2,3}
	got := movesToMakeZigzag(nums)
	fmt.Println(fmt.Sprintf("movesToMakeZigzag result(2) is %v", got))
	time.Sleep(1 * time.Second)

	nums = []int{9,6,1,6,2}
	got = movesToMakeZigzag(nums)
	fmt.Println(fmt.Sprintf("movesToMakeZigzag result(4) is %v", got))
	time.Sleep(1 * time.Second)

	nums = []int{2,7,10,9,8,9}
	got = movesToMakeZigzag(nums)
	fmt.Println(fmt.Sprintf("movesToMakeZigzag result(4) is %v", got))
	time.Sleep(1 * time.Second)

	nums = []int{10,4,4,10,10,6,2,3}
	//n_o := []int{3,4,3,10,9,10,2,3}
	//n_j := []int{10,3,4,3,10,6,7,3}
	got = movesToMakeZigzag(nums)
	fmt.Println(fmt.Sprintf("movesToMakeZigzag result(13) is %v", got))
	time.Sleep(1 * time.Second)

}
