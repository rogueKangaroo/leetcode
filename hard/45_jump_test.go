package hard

import (
	"fmt"
	"testing"
	"time"
)

func Test_jump(t *testing.T) {
	nums := []int{2,3,1,1,4}
	got := jump(nums)
	fmt.Println(fmt.Sprintf("jump result is %v", got))
	time.Sleep(100 * time.Millisecond)
}
