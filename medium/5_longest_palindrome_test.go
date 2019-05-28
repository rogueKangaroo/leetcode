package medium

import (
	"fmt"
	"testing"
	"time"
)

func Test_longestPalindrome(t *testing.T) {
	s := "babad"
	got := longestPalindrome(s)
	fmt.Println(fmt.Sprintf("longestPalindrome result is %v", got))
	time.Sleep(1 * time.Second)
}
