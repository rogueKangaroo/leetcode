package hard

import (
	"fmt"
	"testing"
	"time"
)

func Test_isMatch44(t *testing.T) {
	s := "aa"
	p := "a"
	got := isMatch44(s, p)
	if got != false {
		fmt.Println(fmt.Sprintf("s[%v] p[%v] isMatch44 result[false] is %v", s, p, got))
	}
	time.Sleep(100 * time.Millisecond)

	s = "aa"
	p = "*"
	got = isMatch44(s, p)
	if got != true {
		fmt.Println(fmt.Sprintf("s[%v] p[%v] isMatch44 result[true] is %v", s, p, got))
	}
	time.Sleep(100 * time.Millisecond)

	s = "cb"
	p = "?a"
	got = isMatch44(s, p)
	if got != false {
		fmt.Println(fmt.Sprintf("s[%v] p[%v] isMatch44 result[false] is %v", s, p, got))
	}
	time.Sleep(100 * time.Millisecond)

	s = "adceb"
	p = "*a*b"
	got = isMatch44(s, p)
	if got != true {
		fmt.Println(fmt.Sprintf("s[%v] p[%v] isMatch44 result[true] is %v", s, p, got))
	}
	time.Sleep(100 * time.Millisecond)

	s = "acdcb"
	p = "a*c?b"
	got = isMatch44(s, p)
	if got != false {
		fmt.Println(fmt.Sprintf("s[%v] p[%v] isMatch44 result[false] is %v", s, p, got))
	}
	time.Sleep(100 * time.Millisecond)
}
