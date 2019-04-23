package medium

import (
	"fmt"
	"testing"
	"time"
)

func Test_myAtoi(t *testing.T) {
	str := "42"
	got := myAtoi(str)
	fmt.Println(fmt.Sprintf("myAtoi result is %v", got))
	time.Sleep(1 * time.Second)
}
