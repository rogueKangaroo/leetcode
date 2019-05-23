package simple

import (
	"fmt"
	"testing"
	"time"
)

func Test_maxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	got := maxSubArray(nums)
	fmt.Println(fmt.Sprintf("maxSubArray result is %v", got))
	time.Sleep(100 * time.Millisecond)

	nums = []int{8, -19, 5, -4, 20}
	got = maxSubArray(nums)
	fmt.Println(fmt.Sprintf("maxSubArray result is %v", got))
	time.Sleep(100 * time.Millisecond)
}
