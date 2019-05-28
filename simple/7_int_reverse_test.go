package simple

import (
	"fmt"
	"testing"
	"time"
)

func Test_intReverse(t *testing.T) {
	x := 123
	got := intReverse(x)
	fmt.Println(fmt.Sprintf("intReverse result is %v", got))
	time.Sleep(1 * time.Second)
}
