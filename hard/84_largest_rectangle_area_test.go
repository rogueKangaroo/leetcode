package hard

import (
	"fmt"
	"testing"
	"time"
)

func Test_largestRectangleArea(t *testing.T) {
	x := []int{0}
	got := largestRectangleArea(x)
	fmt.Println(fmt.Sprintf("largestRectangleArea result is %v", got))
	time.Sleep(100 * time.Millisecond)
}
