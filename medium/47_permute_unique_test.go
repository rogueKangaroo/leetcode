package medium

import (
	"fmt"
	"testing"
	"time"
)

func Test_permuteUnique(t *testing.T) {
	nums := []int{1,1,2}
	got := permuteUnique(nums)
	fmt.Println(fmt.Sprintf("permuteUnique result is %v", got))
	time.Sleep(1 * time.Second)
}
