package simple

import (
	"fmt"
	"testing"
	"time"
)

func Test_twoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	got := twoSum(nums, target)
	fmt.Println(fmt.Sprintf("twoSum result is %v", got))
	time.Sleep(1 * time.Second)
}
