package simple

import (
	"fmt"
	"testing"
	"time"
)

func Test_climbStairs(t *testing.T) {
	m := 4
	got := climbStairs(m)
	fmt.Println(fmt.Sprintf("reverseList result is %v", got))
	time.Sleep(1 * time.Second)
}
