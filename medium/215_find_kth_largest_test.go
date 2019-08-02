package medium

import (
	"fmt"
	"testing"
	"time"
)

/**
示例 1:

输入: [3,2,1,5,6,4] 和 k = 2
输出: 5
示例 2:

输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
输出: 4
说明:

**/
func Test_findKthLargest(t *testing.T) {
	nums := []int{3,2,1,5,6,4}
	k := 2
	got := findKthLargest(nums, k)
	fmt.Println(fmt.Sprintf("findKthLargest result is %v", got))
	time.Sleep(100 * time.Millisecond)


	nums = []int{3,2,3,1,2,4,5,5,6}
	k = 4
	got = findKthLargest(nums, k)
	fmt.Println(fmt.Sprintf("findKthLargest result is %v", got))
	time.Sleep(100 * time.Millisecond)

	nums = []int{7,6,5,4,3,2,1}
	k = 5
	got = findKthLargest(nums, k)
	fmt.Println(fmt.Sprintf("findKthLargest result is %v", got))
	time.Sleep(100 * time.Millisecond)
}
