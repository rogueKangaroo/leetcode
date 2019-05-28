package medium

import (
	"fmt"
	"testing"
	"time"
)

func Test_zConvert(t *testing.T) {
	s := "LEETCODEISHIRING"
	numRows := 3
	got := zConvert(s, numRows)
	fmt.Println(fmt.Sprintf("zConvert result is %v", got))
	time.Sleep(1 * time.Second)
}
