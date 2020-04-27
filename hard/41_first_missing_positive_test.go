package hard

import (
	"fmt"
	"testing"
	"time"
)

func Test_firstMissingPositive(t *testing.T) {
	nums := []int{3,4,-1,1}
	got := firstMissingPositive(nums)
	fmt.Println(fmt.Sprintf("firstMissingPositive result is %v", got))
	time.Sleep(1 * time.Second)
}


func Test_firstMissingPositive1(t *testing.T) {
	nums := []int{6,7,8,9,1,2,3,-9,-1,0}
	got := firstMissingPositive(nums)
	fmt.Println(fmt.Sprintf("firstMissingPositive result is %v", got))
	time.Sleep(1 * time.Second)
}