package medium

import (
	"fmt"
	"testing"
	"time"
)

func Test_numDecodings(t *testing.T) {
	s := "226"
	got := numDecodings(s)
	fmt.Println(fmt.Sprintf("numDecodings result is %v", got))
	time.Sleep(1 * time.Second)

	s = "22266"
	got = numDecodings(s)
	fmt.Println(fmt.Sprintf("numDecodings result is %v", got))
	time.Sleep(1 * time.Second)

	s = "12"
	got = numDecodings(s)
	fmt.Println(fmt.Sprintf("numDecodings result is %v", got))
	time.Sleep(1 * time.Second)
}
