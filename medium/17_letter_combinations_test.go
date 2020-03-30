package medium

import (
	"fmt"
	"testing"
	"time"
)

func Test_letterCombinations(t *testing.T) {
	digits := "23"
	got := letterCombinations(digits)
	fmt.Println(fmt.Sprintf("letterCombinations result is %v", got))
	time.Sleep(1 * time.Second)
}
