package medium

import (
	"fmt"
	"testing"
	"time"
)

func Test_generateParenthesis(t *testing.T) {
	n := 4
	got := generateParenthesis(n)
	fmt.Println(fmt.Sprintf("generateParenthesis result is %v", got))
	time.Sleep(1 * time.Second)
}
