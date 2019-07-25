package hard

import (
	"fmt"
	"testing"
	"time"
)

func Test_trap(t *testing.T) {
	height := []int{0,1,0,2,1,0,1,3,2,1,2,1}
	got := trap(height)
	fmt.Println(fmt.Sprintf("trap result is %v", got))
	time.Sleep(100 * time.Millisecond)
}
