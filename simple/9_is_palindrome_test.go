package simple

import (
	"fmt"
	"testing"
	"time"
)

func Test_isPalindrome(t *testing.T) {
	x := -121
	got := isPalindrome(x)
	fmt.Println(fmt.Sprintf("isPalindrome result is %v", got))
	time.Sleep(1 * time.Second)
}
